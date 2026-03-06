package com.memorose.client.exceptions;

public class MemoroseAPIException extends RuntimeException {
    private final int statusCode;
    private final String rawResponse;

    public MemoroseAPIException(int statusCode, String message, String rawResponse) {
        super("API Error " + statusCode + ": " + message);
        this.statusCode = statusCode;
        this.rawResponse = rawResponse;
    }

    public int getStatusCode() {
        return statusCode;
    }

    public String getRawResponse() {
        return rawResponse;
    }
}
