# Envdost
A cli tool (friend) to securely store your config/env files with ease & retrive easily regardless of the work machine. Be at work or home machine sync and update env files with ease with just 1-3 commands at max. Manage your env files with the security of [1password](https://1password.com/).

### Installation
Envdost relies on one password cli for storing files securily and to verify the ownership of the files. So, to continue with the installation you must have a valid 1password account. 
New to 1password ? [👉 click me](https://support.1password.com/explore/get-started/)

#### Pre-requisites
- [1password cli](https://developer.1password.com/docs/cli/v1/get-started/)

>  That's it 😄, we are now good to go.

### Steps:
- Download the latest release of envdost :<br />
  See the release to know more about the latest relase of envdost and [click me 👈](https://github.com/sarojregmi200/envdost/releases/download/v1.0.0/envdost.exe) to download the ` version 1.0.0 ` of envdost.
- Add the envdost to ` path `:<br /><br />
  ``` console
  mkdir c:\envdost\bin\
  mv envdost.exe c:\envdost\bin\ #moving the envdost to a folder envdost inside c drive
  ```
  And then after moving it to your desired location add the location to the path.
  After that you should be able to check your installation by typing ` envdost ` in your terminal or command prompt.

### Usage :
These are the commands available as of the latest version of envdost [v1.0.0](https://github.com/sarojregmi200/envdost/releases/tag/v1.0.0):
- [Signin](https://github.com/sarojregmi200/envdost/edit/main/README.md#signin)
- [Signout](https://github.com/sarojregmi200/envdost/edit/main/README.md#signout)
- [Create](https://github.com/sarojregmi200/envdost/edit/main/README.md#create)
- [Set](https://github.com/sarojregmi200/envdost/edit/main/README.md#set)
- [Push](https://github.com/sarojregmi200/envdost/edit/main/README.md#push)
- [Pull](https://github.com/sarojregmi200/envdost/edit/main/README.md#pull)
- [Delete](https://github.com/sarojregmi200/envdost/edit/main/README.md#delete)
    - Project
    - Env

 #### Signin
 Signin command is used to login to a account, all the logins in this cli tools are based on 1password cli auth so, you will be prompted your way of authentication.
 And after verifying your identity a environment variable will be set with your credentials. 
 <br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost signin
```
<br />

 #### Signout
 Signout command is used to logout from a account. It will fail if you are not logged in, and print a message saying you are not logged in. It clears the environment variables set by the signin command.
 <br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost logout
```
<br />

 #### Create
 Create command is used to create a project in envdost. Where a project is a representation of your current project that contains several number of configuartion files and env or even secret files. You can create a project named todo in a todo project. User must be loggedin to perform this action. This will generate a vault with the provided project name in the authenticated 1password account. You can create multiple projects by separating each name with a space.
 <br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost create [Project_Name1] [Project_Name2]
```
<br />


 #### Set
 Set command is used to select a project as the current project from all the created project. It is case insensitive but must not differ other than case. You must be logged in to use this command if you are not loggedin then you will be prompted to login first, and then after successful signin this command will run. This will set a selected project in the system env variable.
<br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost set [Project_Name1]
```
<br />

 #### Push
 Push command parses the provided file and breaks it into key value pair and stores them effeciently in the under the vault i.e vault created with the ` create ` command. You must be loggedin to use this command. 
 <br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost push [File Name]
```
<br />


 #### Pull
 Pull command fetches and creates the specified config files. If no file name is supplied it fetches all the files that were pushed under the same project. You must be logged in to use this command. You can use ` -r ` flag to fetch the values in reference mode instead of plain text. It can be useful to ensure the highest level of security. 
 <br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost pull [File Name]

⬢ projects ⚡ sample ◉
> envdost pull

⬢ projects ⚡ sample ◉
> envdost pull -r 
```
<br />



 #### Delete
 Delete command is used to delete the projects or the env files. It deletes all the projects will the same name. For now the env subcommand doesnot do anything it will be a feature in next version. 
 <br />
 usage:
 
 ``` console
⬢ projects ⚡ sample ◉
> envdost delete project [Project name] // deletes all the projects with the given name

⬢ projects ⚡ sample ◉
> envdost delete env [file name] // not available in this version.

```
<br />


### Example:
Here is a simple example of a project using envdost to manage the env files. Let us suppose you have a todo project that has some secrets stored in it's ` .env ` file.

```console
todo-app/
├── config/
│   ├── db.config.js
│   └── .env
├── controllers/
│   └── todoController.js
├── models/
│   └── todoModel.js
├── routes/
│   └── todoRoutes.js
├── app.js
└── server.js
```
In the above project the ` .env ` file contains sensitive details like database credentials, different client id, and tokens from thirdparty integrations. Here is how you can use ` envdost ` to securely store it.

 ``` bash
⬢ projects ⚡ todo ◉
> envdost create todo
Enter the password for you@gmail.com at my.1password.com:
Logging in ⣽
Creating todo project  ⢿
Project todo created successfully.


⬢ projects ⚡ todo ◉
> envdost set todo
Selecting todo  ⣾
Project todo is selected successfully


⬢ projects ⚡ todo ◉
> envdost push .\config\.env
Processing file .env
Completed processing .env
Uploading .env ⣟
File .env successfully uploaded under project todo


⬢ projects ⚡ todo ◉
> envdost pull
looking for config files in project todo ⣻
Fetching content of .env ⣯
File .env created successfully.
Completed writing to file .env
```

It is this simple now the contents of you env file are stored in a your 1password account under the vault todo.


### Conclusion
You can manage your env files and important secrets using envdost, You can open a discussion contact me at any of my socials, or create a issue if you are facing any problem. Now you no longer are forced to text yourself the contents of env files over any social media.

### Follow Me

[![Twitter](https://img.shields.io/badge/Twitter-%40sarojregmi200-blue?style=flat&logo=twitter)](https://twitter.com/sarojregmi200)
[![GitHub](https://img.shields.io/badge/GitHub-sarojregmi200-black?style=flat&logo=github)](https://github.com/sarojregmi200)
[![Instagram](https://img.shields.io/badge/Instagram-sarojregmi200-purple?style=flat&logo=instagram)](https://instagram.com/sarojregmi200)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-sarojregmi200-blue?style=flat&logo=linkedin)](https://linkedin.com/in/sarojregmi200)
[![Hashnode](https://img.shields.io/badge/Hashnode-sarojregmi200-yellow?style=flat&logo=hashnode)](https://hashnode.com/@saroj-regmi)
[![Facebook](https://img.shields.io/badge/Facebook-sarojregmi200-blue?style=flat&logo=facebook)](https://facebook.com/sarojregmi200)



Support me by buying me a coffee!

[![Buy Me A Coffee](https://img.shields.io/badge/Buy%20Me%20A%20Coffee-support-%23FFDD00?style=flat&logo=buy-me-a-coffee)](https://buymeacoffee.com/sarojregmi200)
 
