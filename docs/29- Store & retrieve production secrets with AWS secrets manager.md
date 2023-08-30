# Store & retrieve production secrets with AWS secrets manager

```bash
 aws secretsmanager get-secret-value --secret-id g-bank --query SecretString --output text | jq -r 'to_entries |map("\(.key)=\(.value)")|.[]' >> app.env
```
