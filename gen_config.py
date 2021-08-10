import yaml 
defaultData = dict(
    db = dict(
        driver = '',
        connection = ''
    ),
    web_host='localhost',
    web_port='8090'
)
with open('/app/app_config/config.yaml', 'w') as outfile:
    yaml.dump(defaultData, outfile, default_flow_style=False)