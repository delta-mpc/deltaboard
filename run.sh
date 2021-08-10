#!bin/bash
echo $1;
if [ $1 = "init" ];then
python3 gen_config.py;
else
python3 gen_web_config.py;
jupyterhub &
./main migrate --conf config;
./main server --conf config;
fi