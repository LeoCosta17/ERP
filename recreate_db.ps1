$mysqlCmd = "mysql -u admin -p12345 -P 3307 -h 127.0.0.1"
cmd.exe /c "$mysqlCmd -e `"DROP DATABASE IF EXISTS ecommerce; CREATE DATABASE ecommerce;`""
$dbCmd = "$mysqlCmd ecommerce"

# Gestao scripts
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_empresas.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_endereco_empresa.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_credenciais_empresa.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_dados_fiscais_empresa.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_fornecedores.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_enderecos_fornecedor.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_telefones_fornecedor.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_categorias_debito.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_debitos.sql"
cmd.exe /c "$dbCmd < .\Gestao\scripts\cria_tb_usuarios_gestao.sql"

# Ecommerce scripts
cmd.exe /c "$dbCmd < .\Ecommerce\scripts\cria_tb_usuarios.sql"
cmd.exe /c "$dbCmd < .\Ecommerce\scripts\cria_tb_enderecos.sql"
cmd.exe /c "$dbCmd < .\Ecommerce\scripts\cria_tb_codigos_validacao.sql"

Write-Host "Database recreated successfully."
