#!bin/bash
echo $1;
if [ $1 = "init" ];then
python3 gen_config.py;
echo ".open /app/db/delta_dashboard.db" | sqlite3
else
if [ ! -e /app/db/delta_dashboard.db ];then
echo ".open /app/db/delta_dashboard.db" | sqlite3;
fi
python3 gen_web_config.py;
# jupyterhub;
./main migrate --conf /application/config;
jupyterhub &
nginx -g "daemon on;" &
./main server --conf /application/config;
fi