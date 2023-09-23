# Summary

This is an URL shortener service supports 2 APIs

- Submit an URL and generate a corresponding shortened URL
- Access a shorterned URL and redirected to the corresponding URL

## Worklog

1. Start the project and the scaffolding
2. Start with a hello world service

## Pitfalls

1. Scaffolding for docker compose
2. protobuf generation with google defined types
    1. solved by vendoring the import packages
3. gin request binding with protobuf timestamp type
