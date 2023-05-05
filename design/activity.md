# Activity

```mermaid
---
title: Register
---
flowchart TD

  %% Define classes for style
  classDef user fill:#0072BC,stroke:none,color:black
  classDef server fill:#FDB933,stroke:none,color:black

  %% Diagram
  A((START)) --> B(Asks for help):::user
  B --> C(Returns way to send HTTP request to register service):::server
  C --> D(Sends requests to register service):::user
  D --> E{Is it a valid request?}
  E --Yes--> F(Returns API Key):::server
  F --> G(((END)))
  E --No--> H(Makes user to resend a request again):::server
  H --> I(Resend request):::user
  I --> E
```

<hr/>

```mermaid
---
title: Get
---
flowchart TD

  %% Define classes for style
  classDef user fill:#0072BC,stroke:none,color:black
  classDef server fill:#FDB933,stroke:none,color:black
  classDef point fill:black,stroke:none,color:black
```