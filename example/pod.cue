template: {
  output: {
    apiVersion: "v1"
    kind: "Pod"
    spec: {
      containers: [
        {
          image: parameter.image
          ports: [
            {
              containerPort: parameter.port
            }
          ]
        }
      ]
      hostAliases: parameter.hostAliases
    }
  }
  parameter: {
    image: string
    port: int
    hostAliases?: [...{
      ip: string
      hostnames: [...string]
    }]
  }
}
