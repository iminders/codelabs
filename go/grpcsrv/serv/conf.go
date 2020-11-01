package serv

type Conf struct {
	Name      string
}


func LoadConf() (conf Conf, err error) {

    conf = Conf{
        // AccessKey: os.Getenv("MINIO_ACCESS_KEY"),
        // SecretKey: os.Getenv("MINIO_SECRET_KEY"),
        Name: "TestServer",
    }
    return
}
