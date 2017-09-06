# Protobufs Data Validation

This repo outlines the use case for data validation rules inside .proto message definitions. The extension might look as follows:

```
message Example {
    string message = 1; `pattern: $[a-zA-Z].*`
    int32 percentage = 2; `minValue: 0, maxValue: 100`
    repeated int32 list = 3; `minItems: 1, maxItems: 10`
}
```

Without validation rules in the .proto definitions, the developer is required to maintain boilerplate validation code. This repo implements an example of this in the `example/validate.go` file. This is a major pain because code must be maintained for each language that is expected to work with the Example message and must be maintained consistently.

This validate code is completely boilerplate, and it could be so much better if these type of schema rules were supported in the .proto definition itself. The validation code could be generated and baked into the proto.Marshal and proto.Unmarshal code in an automated and consistent fashion.