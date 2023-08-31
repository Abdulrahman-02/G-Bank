# Generate DB documentation page and schema SQL dump from DBML

```bash
npm install -g dbdocs
# Generate DB documentation page
dbdocs build  doc/db.dbml
# Set up a password
dbdocs password --set password --project project-name
# Install dbml cli
npm install -g @dbml/cli
# Generate SQL schema dump
dbml2sql --postgres -o doc/schema.sql doc/db.dbml
```
