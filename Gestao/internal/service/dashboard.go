package service

import (
	"context"
	"gestao/internal/repository"
	"sync"
	"time"
)

type DashboardService struct {
	Repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{Repo: repo}
}

type ResumoDashboard struct {
	TotalVencido     float64                     `json:"total_vencido"`
	TotalSemana      float64                     `json:"total_semana"`
	VendaDia         float64                     `json:"venda_dia"`
	VendaMes         float64                     `json:"venda_mes"`
	DespesasCategoria []repository.CategoriaGasto `json:"despesas_categoria"`
}

// ObterResumo consulta e agrega todos os dados do dashboard em uma única chamada usando goroutines para performance
func (s *DashboardService) ObterResumo(ctx context.Context) (*ResumoDashboard, error) {
	hoje := time.Now()

	// Lógica para encontrar a Segunda-feira (início da semana) e Domingo (fim da semana)
	// Considerando que a semana começa na Segunda-feira
	offsetSegunda := int(time.Monday - hoje.Weekday())
	if offsetSegunda > 0 {
		offsetSegunda = -6
	}
	inicioSemana := hoje.AddDate(0, 0, offsetSegunda)
	fimSemana := inicioSemana.AddDate(0, 0, 6)

	// Lógica para início e fim do mês atual
	inicioMes := time.Date(hoje.Year(), hoje.Month(), 1, 0, 0, 0, 0, hoje.Location())
	fimMes := inicioMes.AddDate(0, 1, -1)

	resumo := &ResumoDashboard{}
	var errGeral error
	var wg sync.WaitGroup

	// Executa as queries em paralelo
	wg.Add(3)

	go func() {
		defer wg.Done()
		total, err := s.Repo.GetTotalDebitosAtrasados(ctx)
		if err != nil {
			errGeral = err
			return
		}
		resumo.TotalVencido = total
	}()

	go func() {
		defer wg.Done()
		total, err := s.Repo.GetTotalDebitosSemana(ctx, inicioSemana, fimSemana)
		if err != nil {
			errGeral = err
			return
		}
		resumo.TotalSemana = total
	}()

	go func() {
		defer wg.Done()
		cats, err := s.Repo.GetDespesasPorCategoria(ctx, inicioMes, fimMes)
		if err != nil {
			errGeral = err
			return
		}
		resumo.DespesasCategoria = cats
	}()

	wg.Wait()

	if errGeral != nil {
		return nil, errGeral
	}

	// TODO: Quando o módulo de Vendas for criado, substituir os mocks abaixo por chamadas ao repository
	resumo.VendaDia = 0.0 // Mock provisório
	resumo.VendaMes = 0.0 // Mock provisório

	return resumo, nil
}
