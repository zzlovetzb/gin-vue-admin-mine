mongo:
    docker run -d --name mongodb -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -p 27017:27017 mongo:4.4.3
postgre:
    docker run --name postgres -e POSTGRES_PASSWORD=root -p 5438:5432 -d postgres
mysql:
    cd /media/zhgl/inst_96/mysql-5.7.42
    cd support-files
    ./mysql.server start
    cd ..
    cd bin
    sudo ./mysql -u root -p
redis:
    cd redis-6.0.6
    make
    cd src
    redis-server
