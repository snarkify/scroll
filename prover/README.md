# Run batch prover
## How-to run Snarkify batch prover docker
under the scroll/prover directory, run
1. `git switch snarkify-gpu-deploy`
2. `make snarkify`
3. `docker build -t scroll-prover-gpu:latest .`
4. `docker run -v /home/ubuntu/scroll/volume:/snarkify-data  -p 8080:8080 scroll-prover-gpu`
5. In a different shell, run `run_batch.sh` to submit a sample job to the prover to generate proof

## How-to run Snarkify batch prover remotely
ssh to gpu-5
1. `cd scroll/scroll/prover`
2. build the docker image following above instructions
3. `snarkify deploy --tag "{your_tag}" --image scroll-prover-gpu:latest`
4. Follow the printed instruction to check if the deployment is done, it should show {your_tag} if it is successful
5. `snarkify task create --file prover_input.json` to create a new proof task
6. you should get a {task_id} if the task is created, then use `snarkify task log {task_id}` to stream the logs

# Run chunk prover
## How-to run Snarkify chunk prover docker
under the scroll/prover directory, run
1. `git switch chunk_prover`
2. `make snarkify`
3. `docker build -t scroll-chunk-prover:latest .`
4. `docker run -v /home/ubuntu/scroll/volume:/snarkify-data  -p 8080:8080 scroll-chunk-prover`
5. In a different shell, run `run_real_chunk.sh` to submit a sample job to the prover to generate proof
