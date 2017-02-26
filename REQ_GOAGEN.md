# Package call flows in goagen

| goagen command | generator packages | writer packages | generated packages |
| --- | --- | --- | --- |
| `server` | codegen/generators/server | codegen/writers/endpoint | gen/endpoints |
| | | codegen/writers/service | gen/services |
| | | rest/codegen | gen/transport/http |
| `client` | codegen/generators/client | codegen/writers/endpoint | gen/endpoints |
| | | codegen/writers/service | gen/services |
| | | rest/codegen | gen/transport/http |
| | | | gen/cli |
| `swagger` | codegen/generators/swagger |
