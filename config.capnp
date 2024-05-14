using Workerd = import "/workerd/workerd.capnp";

const config :Workerd.Config = (
  services = [ (name = "main", worker = .worker) ],
  sockets = [ ( name = "http", address = "*:8080", http = (), service = "main" ) ]
);

const worker :Workerd.Worker = (
  modules = [
    ( name = "entrypoint", esModule = embed "./build/entrypoint.mjs" ),
    ( name = "./app.wasm", wasm = embed "./build/app.wasm" )
  ],
  compatibilityDate = "2023-03-14",
);
