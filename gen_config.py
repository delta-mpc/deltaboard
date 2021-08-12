import yaml 
import os
os.makedirs("/app/app_config",exist_ok=True)
os.makedirs("/app/db",exist_ok=True)
defaultData = dict(
    db = dict(
        driver = '',
        connection = ''
    ),
    node_address ='localhost',
    port='8090'
)
with open('/app/app_config/config.yaml', 'w') as outfile:
    yaml.dump(defaultData, outfile, default_flow_style=False)