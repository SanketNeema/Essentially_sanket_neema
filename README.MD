# Backend Task | GOLANG | Essentially AI

Being a firm in the capital markets of India, we have to work with huge amounts of data. You can think of a data packet flowing into the system every nanosecond.

Working with this high frequency of data needs our systems to be extremely robust and error tolerant while keeping the code maintainable. This led to our choice of Golang as the backend.

Make sure the code which you're writing follows these principles.

## Task: DECODE!!!

The source of our high frequency data gives us an encoded data in the form of a byte stream.

The task is to make a function to decode the packet using the decoding format given below and return the decoded message in the form of a struct. The packet size is fixed at 44 bytes.

Decoding map:
1. First 2 bytes represent a short datatype
2. The next 12 bytes represent 12 characters
3. The next 1 byte represent a single byte
4. The next 8 bytes represent 8 characters
5. The next 2 bytes represent a short datatype
6. The next 15 bytes represent 15 characters
7. The next 4 bytes represent a long datatype

Below is a sample test case:

Packet: [\x04 \xD2 \x6B \x65 \x65 \x70 \x64 \x65 \x63 \x6F \x64 \x69 \x6E \x67 \x38 \x64 \x6F \x6E \x74 \x73 \x74 \x6F \x70 \x03 \x15 \x63 \x6F \x6E \x67 \x72 \x61 \x74 \x75 \x6C \x61 \x74 \x69 \x6F \x6E \x73 \x07 \x5B \xCD \x15]

Decoded struct: {1234, "keepdecoding", 56, "dontstop", 789, "congratulations", 123456789}