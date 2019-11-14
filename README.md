# Simple Locker Application with Command Line

Simple Locker Application with Command Line

## Requirements
- [Golang](https://golang.org/dl/)

## Installation Guide

- Before installing this Project, make sure your computer already installed application in the `requirements` First.

- Clone this Project to your local Computer

## Usage
- Run the program
    ```
    go run main.go
    ```
- Initial size of locker using. example :
    ```
    init [size of locker]
    ```
- Input the card identity with type Identity and Identity Number. example : 
    ```
    input [type id] [id number]
    ```
- Check the list of locker status.
    ```
    status
    ```
- For empty the locker use number of locker you want. example :
    ```
    leave [number of locker]
    ```
- To find the Locker number by identity number. example :
    ```
    find [id number]
    ```
- To search all list of identity number by type of identity. example : 
    ```
    search [type id]
    ```
- For exit the program.
    ```
    exit
    ```

## Example
Below is example to run the command line. 
    ```
    go run main.go
    init 5
    input SIM 12345
    status
    leave 2
    find 12345
    search SIM
    exit
    ```
## Run the test
To run the test use this command. 
    ```
    go test
    ```