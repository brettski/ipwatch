

mod env_config;
mod http_requests;

fn main() {
    println!("starting up ipWatch\n\n");
    println!("Reading .env file");
    dotenv::dotenv().ok();
    
    let config = env_config::create();

    println!("our config {:#?}", config);

    http_requests::fetch_ip(&config.endpoint);


}


