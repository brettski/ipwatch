use std::env;

#[derive(Debug)]
pub struct EnvConfig {
    pub endpoint: String,
    pub test_location: String,
    pub postmark: PostmarkEnv,
}

#[derive(Debug)]
pub struct PostmarkEnv {
    pub token: String,
    pub email_to: String,
    pub email_from: String,
}

pub fn create() -> EnvConfig {
    // load up new struct
    // PostmarkEnv first
    let token: String;
    match env::var("POSTMARK_TOKEN") {
        Ok(val) => token = val,
        Err(e) => panic!("error getting POSTMARK_TOKEN env value, {e}"),
    };
    let email_to: String;
    match env::var("EMAIL_TO") {
        Ok(val) => email_to = val,
        Err(e) => panic!("error getting EMAIL_TO env value, {e}"),
    };
    let email_from: String;
    match env::var("EMAIL_FROM") {
        Ok(val) => email_from = val,
        Err(e) => panic!("error getting EMAIL_FROM env value, {e}"),
    };
    let postmark = PostmarkEnv {
        token,
        email_to,
        email_from,
    };


    //let mut EnvConfig: new_envconfig;
    let endpoint: String;
    match env::var("ENDPOINT_CHK") {
        Ok(val) => endpoint = val,
        Err(e) => panic!("error getting ENDPOINT_CHK env value, {e}"),
    };

    let test_location: String;
    match env::var("TEST_LOCATION") {
        Ok(val) => test_location = val,
        Err(e) => panic!("error getting ENDPOINT_CHK env value, {e}"),
    };

    let env_config = EnvConfig {
        endpoint,
        test_location,
        postmark
    };

    env_config
       
}
    
