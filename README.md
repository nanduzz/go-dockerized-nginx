# go-dockerized-nginx

## How to run 

if is the first time you iwll run the project then run command on the root of the project:

    $ docker-compose up -d --build 

the --build is needed only when you don't have the images on your computer or in cases that you did some modifications. after this you can run

If you alread have the images you can simple run:

    $ docker-compose up -d 

Then in your web browser try to access http://localhost:8080, every time you refresh the page nginx will send your request to a different app and you should see a different number. 

If you want do stop all the containers of the docker-compose file you can use the following command in the root dir of the project

    $ docker-compose down
