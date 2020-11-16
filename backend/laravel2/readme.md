<p align="center"><img src="https://laravel.com/assets/img/components/logo-laravel.svg"></p>


## Laravel Simple API CRUD and follow below instruction

## Installation Instruction

- Clone the Repo.
- Run 'composer install'
- Run 'cd apicrud'
- Run 'cp .env.example .env' and change database information
- Run 'php artisan migrate'

## Insert Data

- Insert student data into database using postman.

- Method: POST, URL: http://127.0.0.1:8000/api/student

<img src="https://github.com/Shobuj718/apicrud/blob/master/public/image/insert%20data.PNG">

## Show All Data

- Show all student data from database using postman .

- Method: GET, URL: http://127.0.0.1:8000/api/students

<img src="https://github.com/Shobuj718/apicrud/blob/master/public/image/show%20all%20data.PNG">

## Show Single Data

- Show Single student data from database using postman .
- Method: GET, URL: http://127.0.0.1:8000/api/student/1

<img src="https://github.com/Shobuj718/apicrud/blob/master/public/image/show%20single%20data.PNG">

## Update Student Data

- Update Student data into database using postman .

- Method: PUT, URL: http://127.0.0.1:8000/api/student/1?fname=yourfirstname&lname=yourlastname&email=youremail&password=anypassword

<img src="https://github.com/Shobuj718/apicrud/blob/master/public/image/updata%20data.PNG">

## Delete Student Data

- Delete Student data into database using postman .

- Method: DELETE, URL: http://127.0.0.1:8000/api/student/1

<img src="https://github.com/Shobuj718/apicrud/blob/master/public/image/delete%20data.PNG">
