# Protocol Buffers #

gRPC uses protocol buffers for communications. This buffers have smaller content size, for example:

JSON: 52 bytes (compressed)
```
{
    "age": 26,
    "first_name": "Clement",
    "last_name": "JEAN"
}
```

gRPC Protocol Buffer: 17 bytes
```
syntax = "proto3"

message Person {
    uint32 age = 1
    string first_name = 2
    string last_name = 3
}
```

This saves bandwidth and storage space because messages are smaller. Also, being serialized, it requires less CPU.

# About gRPC #

gRPC relies on HTTP2 which is much faster than HTTP1.1. HTTP2 allows us to have one single TCP connection instead of several tcp connections that need to be established in the version 1.1.

# Types of API in gRPC #

We can have 4 types of APIs in gRPC

## Unary API

Client will send one reqeust and wait for a server response

```
rpc Greet(GreetRequest) returns (GreetResponse) {};
```

## Server Streaming

Client will send a request and the server might request one or more responses

```
rpc GreetManyTimes(GreetRequest) returns (stream GreetResponse) {};
```

## Client Streaming

The client can send one or more requests and the server returns one response

```
rpc LongGreet(stream GreetRequest) returns (GreetResponse) {};
```

## Bi Directional Stremaming

The client can send multiple requests and also the server can respond with several requests
```
rpc GreetEveryone(stream GreetRequest) returns (stream GreetResponse) {};
```

# Scalability in gRPC

The server is async, which means that the main thread is not locked in one single request and can answer multiple requests at the same time.
On the client side, we can use async or sync. For example, when a response is critical, we might need sync responses.

# Security in gRPC

Schema based serialization. This means that the data is binary and, therefore, not human readable. In top of this, there is easy to use SSL certificates.

# gRPC vs REST

Protocol Buffers vs JSON
HTTP2 vs HTTP1
Streaming vs Unary
Bi Directional vs Client -> Server
Free Design vs GET/POST/UPDATE/DELETE
