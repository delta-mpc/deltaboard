import yaml 
import os
os.makedirs("/app/app_config",exist_ok=True)
os.makedirs("/app/db",exist_ok=True)
defaultData = dict(
    db = dict(
        driver = '',
        connection = ''
    ),
    node_address ='http://localhost',
    port='8090',
    public_registration= True,
    post_registry_url= "/post-regist.html"
)
with open('/app/app_config/config.yaml', 'w') as outfile:
    yaml.dump(defaultData, outfile, default_flow_style=False)