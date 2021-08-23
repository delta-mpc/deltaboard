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
    allow_origins = ['http://localhost:8090','https://board.deltampc.com'],
    node_address='http://localhost',
    public_registration= True,
    post_registry_url= "/post-regist.html",
    host='localhost',
    port='8090'
)
yamlconfig = {}

with open("/app/app_config/config.yaml", "r") as f:
   yamlconfig = yaml.full_load(f)
with open("/application/config/config.yaml","w") as f:
   defaultData['allow_origins'] = ['http://localhost:8090']
   defaultData['db'] = yamlconfig['db']
   defaultData['node_address'] = yamlconfig['node_address']
   defaultData['public_registration'] = yamlconfig['public_registration']
   defaultData['post_registry_url'] = yamlconfig['post_registry_url']
   yaml.dump(defaultData, f, default_flow_style=False)
with open("/application/web/config.json","w") as cf:
   json.dump({"port":8090},cf)
with open("/application/web/.env.development","w") as devFile:
   devFile.write("VUE_APP_BASE_API=http://localhost:8090")
