import yaml 
import os
if not os.path.exists("/app/db"):
   os.makedirs("/app/db",exist_ok=True)
if not os.path.exists('/app/app_config/config.yaml'):
   os.makedirs("/app/app_config",exist_ok=True)
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