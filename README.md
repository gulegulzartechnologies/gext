# gext
Manage a GITHUB folder

# Commands

- `gext init` 
    - check presence of `~/GITHUB` folder 
    - if present, display `GITHUB folder exists at $HOME`
    - if not present, create `~/GITHUB` folder
- `gext list`
    - `gext list users` - list all users
    - `gext list repos <user>` - list all repositories cloned under `<user>`
- `gext clone <URL>`
    - check if URL is a github url and correctly formatted
    - check if `<user>` has folder in `~/GITHUB`
    - clone repo in `<user>` folder
    