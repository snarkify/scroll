use std::io;

use async_trait::async_trait;
use snarkify_sdk::prover::ProofHandler;
use prover_runner::{
    prover_core::Prover,
    types::{Task, ProofDetail},
    config::{Config, AssetsDirEnvConfig},
    version,
};
use std::cell::RefCell;

struct MyProofHandler;


fn init_prover() -> Prover<'static> {
    let config: Config = Config::from_file("config.json".to_string()).expect("Failed to load config");

    if let Err(e) = AssetsDirEnvConfig::init() {
        log::error!("AssetsDirEnvConfig init failed: {:#}", e);
        std::process::exit(-2);
    }

    log::info!(
        "Starting prover. name: {}, type: {:?}, version: {}",
        config.prover_name,
        config.proof_type,
        version::get_version(),
    );

    Prover::new(Box::leak(Box::new(config))).expect("Failed to create prover")
}


thread_local! {
    static PROVER: RefCell<Prover<'static>> = RefCell::new(init_prover());
}


#[async_trait]
impl ProofHandler for MyProofHandler {
    type Input = Task;
    type Output = ProofDetail;
    type Error = String;

    async fn prove(data: Self::Input) -> Result<Self::Output, Self::Error> {
        PROVER.with_borrow(|p| {
            p.prove_task(&data).map_err(|e| e.to_string())
        })
    }
}

fn main() -> Result<(), io::Error> {
    snarkify_sdk::run::<MyProofHandler>()
}