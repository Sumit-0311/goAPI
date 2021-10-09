<h1>Simple API implemented using Golang</h1>

# Get Started

Make sure you have Golang and MongoDB both installed on your system.
Once you've insured these requirements are met, download or clone this repo on your own machine
and execute the following commands in order:

```
go install
```

<p>Clone the repository using the following URL</p>

```
https://github.com/Sumit-0311/goAPI
```

One this command is successfully executed, you're good to go!

To start the process, input the following command in terminal:

```
go run .
```

Once, you run this command, you install all the dependencies, and then you should get the following output.

```
Starting the application...
```

Now the API is running on your machine on the local port 12345, next lets move onto testing out newly
implemented API!

# Test

It is recommended use Postman to input data into the dataset and compass to monitor your input data.
Input data in the formats of the sturctures provided in posts.go and user.go in the repo for posts and
users respectively.
To input data in user collection the link is: http://localhost/12345/user
To input data in post collection the link is: http://localhost/12345/post
To get user data out of user collection the link is: http://localhost/12345/user/{id} replace id by the user's ID you want.
To all data in user collection the link is: http://localhost/12345/users
To get post data out of user collection the link is: http://localhost/12345/posts/users/{id} replace id by the user's ID whose post you want to see.
