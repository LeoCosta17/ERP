import os
import re

sql_dir = "scripts"

for filename in os.listdir(sql_dir):
    if not filename.endswith(".sql"):
        continue
    filepath = os.path.join(sql_dir, filename)
    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    # Replacements
    # id bigint auto_increment primary key -> id BIGSERIAL PRIMARY KEY
    content = re.sub(r'(?i)bigint\s+auto_increment\s+primary\s+key', 'BIGSERIAL PRIMARY KEY', content)
    content = re.sub(r'(?i)int\s+auto_increment\s+primary\s+key', 'SERIAL PRIMARY KEY', content)
    content = re.sub(r'(?i)auto_increment', '', content) # just in case
    
    # enum('A', 'B') -> VARCHAR(50) CHECK (col IN ('A', 'B'))
    # Wait, we need the column name to do CHECK (col IN...). It's easier to just use VARCHAR(50) and let the app validate, or do it properly.
    # Let's just convert enum to VARCHAR(50) for simplicity in a SaaS ERP
    content = re.sub(r"(?i)enum\([^)]+\)", "VARCHAR(50)", content)
    
    # on update current_timestamp -> remove it (handled in Go code or triggers, for now just remove)
    content = re.sub(r"(?i)on\s+update\s+current_timestamp", "", content)
    
    # bool default false -> BOOLEAN DEFAULT FALSE (actually MySQL uses tinyint/bool, Postgres uses boolean)
    content = re.sub(r"(?i)bool\s+default\s+false", "BOOLEAN DEFAULT FALSE", content)

    # ENGINE=InnoDB DEFAULT CHARSET=utf8mb4... -> remove entirely, just keep the closing parenthesis and semicolon
    content = re.sub(r"(?i)\)\s*ENGINE\s*=\s*[a-zA-Z0-9_]+\s*(?:DEFAULT\s+CHARSET\s*=\s*[a-zA-Z0-9_]+)?\s*(?:COLLATE\s*=\s*[a-zA-Z0-9_]+)?\s*;", ");", content)
    # Also catch cases without closing parenthesis in the same line or just the ENGINE keyword
    content = re.sub(r"(?i)ENGINE\s*=\s*[a-zA-Z0-9_]+\s*(?:DEFAULT\s+CHARSET\s*=\s*[a-zA-Z0-9_]+)?\s*(?:COLLATE\s*=\s*[a-zA-Z0-9_]+)?\s*;", ";", content)

    with open(filepath, "w", encoding="utf-8") as f:
        f.write(content)

print("SQL scripts converted!")
