
Enable general logging

```yaml
---
AWSTemplateFormatVersion: "2010-09-09"
Description: A sample template
AWSTemplateFormatVersion: 2010-09-09
Description: Good example
Resources:
  Broker:
    Type: AWS::AmazonMQ::Broker
    Properties:
      Logs:
        General: true
```