import json
import yaml
defaultData = dict(
    environment = 'debug',
    db = dict(
        driver = '',
        connection = ''
    ),
    log = dict(
       level = 'debug',
    ),
    http = dict(
       host = '0.0.0.0',
       port = '8080'
    ),
    session = dict(
       driver = "memory",
       secret_key = "secret_key",
       connection = "",
       login_max_age_seconds = 36000,
       vault_max_age_seconds = 3600
    ),
    allow_origins = ['http://localhost:8090'],
    web_host='localhost',
    web_port='8090'
)
yamlconfig = {}

with open("/app/app_config/config.yaml", "r") as f:
   yamlconfig = yaml.full_load(f)
with open("/app/config/config.yaml","w") as f:
   defaultData['web_port'] = yamlconfig['web_port']
   defaultData['allow_origins'] = ['http://localhost:' + yamlconfig['web_port']]
   defaultData['db'] = yamlconfig['db']
   yaml.dump(defaultData, f, default_flow_style=False)

with open("/app/web/config.json","w") as cf:
   json.dump({"port":yamlconfig['web_port']},cf)

with open("/app/web/.env.development","w") as devFile:
   devFile.write("VUE_APP_BASE_API=http://localhost:" + yamlconfig['web_port'])
 
