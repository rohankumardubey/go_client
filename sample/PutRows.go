package main

import(
    "github.com/griddb/go_client"
    "fmt"
    "os"
    "strconv"
)

func main() {
    factory := griddb_go.StoreFactoryGetInstance()
    defer griddb_go.DeleteStoreFactory(factory)
    containerName := "SampleGo_PutRows"

    // Get GridStore object
    port, err := strconv.Atoi(os.Args[2])
    if (err != nil) {
        fmt.Println("strconv port failed", err)
        os.Exit(2)
    }
    gridstore, err := factory.GetStore(map[string]interface{} {
        "host": os.Args[1],
        "port": port,
        "cluster_name": os.Args[3],
        "username": os.Args[4],
        "password": os.Args[5]})
    fmt.Println("Connect to Cluster")
    if (err != nil) {
        fmt.Println("Get Store failed", err)
        panic("err get store")
    }
    defer griddb_go.DeleteStore(gridstore)

    //Create Collection
    conInfo, err := griddb_go.CreateContainerInfo(map[string]interface{} {
        "name": containerName,
        "column_info_list":[][]interface{}{
            {"id", griddb_go.TYPE_INTEGER},
            {"productName", griddb_go.TYPE_STRING},
            {"count", griddb_go.TYPE_INTEGER}},
        "type": griddb_go.CONTAINER_COLLECTION,
        "row_key": true})
    if (err != nil) {
        fmt.Println("Create containerInfo failed, err:", err)
        panic("err CreateContainerInfo")
    }
    defer griddb_go.DeleteContainerInfo(conInfo)
    gridstore.DropContainer(containerName)

    //Create Collection
    col, err := gridstore.PutContainer(conInfo)
    if (err != nil) {
        fmt.Println("put container failed, err:", err)
        panic("err PutContainer")
    }
    defer griddb_go.DeleteContainer(col)
    fmt.Println("Create Collection name", containerName)

    //Register multiple rows
    //(1)Get the container
    col, err1 := gridstore.GetContainer(containerName)
    if (err1 != nil) {
        fmt.Println("GetContainer failed, err:", err1)
        panic("err GetContainer")
    }
    defer griddb_go.DeleteContainer(col)

    //(2) Register multiple rows
    rowList := [][]interface{}{}
    for i := 0; i < 5; i++ {
        rowList = append(rowList, []interface{}{i, "dvd", 10*i})
    }

    err2 := col.MultiPut(rowList)
    if err2 != nil {
        fmt.Println("MultiPut err: ", err2)
    }
    fmt.Println("Put Rows")
    fmt.Println("success!")
}