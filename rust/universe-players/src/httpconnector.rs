use http::{Request, Response};
use curl::easy::Easy;
use std::io::Write;

pub fn make_Request(url: &str) {
    let mut easy = Easy::new();
    easy.url(url).unwrap();
    easy.write_function(|data| {
        Ok(std::io::stdout().write(data).unwrap())
    }).unwrap();
    let content = easy.perform().unwrap();
    println!("{:?}", content);
}


pub fn request_Message (url: &str) {
    
    //concatenate this string with the url
    let request = format!("GET / HTTP/1.1\r\nHost: {}\r\n\r\n", url);
    println!("{}", request);
}