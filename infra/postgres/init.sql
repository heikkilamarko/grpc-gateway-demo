CREATE DATABASE grafana;

CREATE DATABASE grpc_gateway_demo;

\c grpc_gateway_demo;

CREATE SCHEMA demo;

CREATE TABLE demo.success_requests (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    data JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE demo.error_requests (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    error_message TEXT,
    data JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
