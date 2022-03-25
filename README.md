# Details
To see the project endpoint layout or to test it please check  `postman_collection.json`

## Project Structure

### Example Sequence Diagram For 
```mermaid
sequenceDiagram
HTTP ->> Controller Layer: Information
Controller Layer ->> Application Layer: Passes data according to protocols
Application Layer ->> Models: Starts the communication with 
Models->> Repository: Applies Business Logic 
Repository ->> Models: Gives Raw Data
Models ->> Application Layer: Here is data with the way you want
Application Layer ->> Controller Layer: Pass the data
Controller Layer ->> HTTP: Here is the data you wanted
```

## Steps

### Clone the project
`git clone https://github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-oguzhantasimaz.git`

### To run the project
Change directory into project folder

### Change MySQLDB connection information from code
Change MYSQLDB connection information which is in main.go, line 23 and 24

### Start start.sh script
Write `./start.sh` code snippet to the terminal / command-line.
If you get "Permission denied" error just type this to terminal / command-line
`chmod u+r+x start.sh`

#### If sh script doesn`t work
Change directory to
`cd /cmd/homework-4-oguzhantasimaz`

##### Start the program by typing
`go run .`


