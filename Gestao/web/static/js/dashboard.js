document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/';
        return;
    }

    const formatCurrency = (value) => new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(value);

    try {
        const res = await fetch('/api/dashboard/resumo', {
            headers: { 'Authorization': `Bearer ${token}` }
        });

        if (res.status === 401) {
            window.location.href = '/';
            return;
        }

        if (!res.ok) {
            console.error('Erro ao buscar dados do dashboard');
            return;
        }

        const data = await res.json();

        // Popula os KPIs
        document.getElementById('kpi-semana').textContent = formatCurrency(data.total_semana);
        document.getElementById('kpi-vencido').textContent = formatCurrency(data.total_vencido);
        document.getElementById('kpi-venda-dia').textContent = formatCurrency(data.venda_dia);
        document.getElementById('kpi-venda-mes').textContent = formatCurrency(data.venda_mes);

        // Renderiza o gráfico de pizza
        const categorias = data.despesas_categoria || [];
        if (categorias.length === 0) {
            document.getElementById('graficoCategorias').classList.add('d-none');
            document.getElementById('grafico-vazio').classList.remove('d-none');
            return;
        }

        const labels = categorias.map(c => c.categoria);
        const values = categorias.map(c => c.total);
        
        // Cores vibrantes e harmônicas
        const backgroundColors = [
            '#4e73df', '#1cc88a', '#36b9cc', '#f6c23e', '#e74a3b', '#858796'
        ];

        const ctx = document.getElementById('graficoCategorias').getContext('2d');
        new Chart(ctx, {
            type: 'doughnut',
            data: {
                labels: labels,
                datasets: [{
                    data: values,
                    backgroundColor: backgroundColors,
                    hoverBackgroundColor: backgroundColors,
                    hoverBorderColor: "rgba(234, 236, 244, 1)",
                }],
            },
            options: {
                maintainAspectRatio: false,
                plugins: {
                    legend: {
                        position: 'bottom',
                        labels: {
                            padding: 20,
                            font: {
                                family: 'Inter, sans-serif'
                            }
                        }
                    },
                    tooltip: {
                        callbacks: {
                            label: function(context) {
                                let label = context.label || '';
                                if (label) {
                                    label += ': ';
                                }
                                if (context.parsed !== null) {
                                    label += formatCurrency(context.parsed);
                                }
                                return label;
                            }
                        }
                    }
                },
                cutout: '70%',
            },
        });

    } catch (err) {
        console.error('Erro interno:', err);
    }
});
