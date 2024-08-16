use std::{
    env::args,
    fs::File,
    io::{BufRead, BufReader, Read, Write},
    net::TcpStream,
    path::Path,
};

fn main() {
    let addr = args().nth(1).unwrap();
    println!("SERVER ADDRESS IS {}", addr);

    let path = args().nth(2).unwrap();
    println!("PATH IS {}", path);

    let mut stream = TcpStream::connect(format!("{}:1919", addr)).unwrap();
    println!("CONNECTED TO {}", stream.peer_addr().unwrap());

    let mut file = File::open(path.clone()).unwrap();
    let filename = Path::new(path.as_str())
        .file_name()
        .unwrap()
        .to_str()
        .unwrap();
    println!("FILE {} OPENED", filename);

    stream
        .write(format!("{}\n", filename).as_str().as_bytes())
        .unwrap();
    println!("FILENAME SENT");

    let mut br = BufReader::new(stream.try_clone().unwrap());
    let mut msg = String::new();
    br.read_line(&mut msg).unwrap();
    msg = msg.trim().to_string();
    println!("{} RECEIVED FROM THE SERVER", msg);

    let mut bytes = 0;
    let mut buf: [u8; 8192] = [0; 8192];

    loop {
        let newbytes = file.read(&mut buf).unwrap();
        if newbytes == 0 {
            break;
        }
        stream.write(buf.split_at(newbytes).0).unwrap();
        bytes += newbytes;
    }

    println!("SENT {} BYTES", bytes);
    stream.shutdown(std::net::Shutdown::Both).unwrap();
    println!("DONE");
}
