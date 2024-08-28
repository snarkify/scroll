use crate::types::CommonHash;
use anyhow::Result;
use ethers_core::types::BlockNumber;
use tokio::runtime::Runtime;

use serde::{de::DeserializeOwned, Serialize};
use std::fmt::Debug;
use std::thread;
use ethers_providers::{Http, Provider};

use tokio::runtime::Builder;

pub struct GethClient {
    id: String,
    provider: Provider<Http>,
    rt: Runtime,
}

impl GethClient {
    pub fn new(id: &str, api_url: &str) -> Result<Self> {
        let provider = Provider::<Http>::try_from(api_url)?;
        let rt = tokio::runtime::Builder::new_current_thread()
            .enable_all()
            .build()?;

        Ok(Self {
            id: id.to_string(),
            provider,
            rt,
        })
    }

    pub fn get_block_trace_by_hash<T>(&mut self, hash: &CommonHash) -> Result<T>
    where
        T: Serialize + DeserializeOwned + Debug + Send + 'static,
    {
        log::info!(
            "{}: calling get_block_trace_by_hash, hash: {:#?}",
            self.id,
            hash
        );

        // Clone the provider to move it into the thread safely
        let provider_clone = self.provider.clone();
        let hash_clone = hash.clone();  // Clone the hash to move it into the thread

        // Spawn a new thread to run the async code
        let result = thread::spawn(move || {
            // Create a new Tokio runtime using the Builder
            let rt = Builder::new_current_thread()
                .enable_all()
                .build()
                .expect("Failed to create Tokio runtime");

            // Block on the async operation within the blocking thread
            rt.block_on(async {
                let trace_future = provider_clone.request("scroll_getBlockTraceByNumberOrHash", [format!("{hash_clone:#x}")]);
                trace_future.await.map_err(|e| anyhow::Error::new(e))
            })
        })
            .join()
            .expect("Thread panicked");

        result
    }

    pub fn block_number(&mut self) -> Result<BlockNumber> {
        log::info!("{}: calling block_number", self.id);

        let trace_future = self.provider.request("eth_blockNumber", ());

        let trace = self.rt.block_on(trace_future)?;
        Ok(trace)
    }
}
