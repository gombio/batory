# Example DB structure (YAML)

accounts:
  gombio: #TODO: Use as host? Like account name in AWS or Slack, ex. gombio.batory.io
    owner: #TODO: SuperAdmin / Owner account - there can be only one, treat as root
      name: "John Doe"
      email: "john@example.com"
      phone: "+123123123"
    users:
      john@example.com:
        name: "John Doe"
        email: "john@example.com" #TODO: can be used as login
        phone: "+123123123" #TODO: can be used as login
        password: "foobar"
        permissions: [] #TODO: array of permissions, ex. "projects_manager, users_manager", etc.
    people:
      john@example.com: #TODO: Can be transfered to "users", so that a person can log in
        name: "John Doe"
        email: "john@example.com"
        phone: "+123123123"
        jane@example.com:
          name: "Jane Doe"
          email: "jane@example.com"
          phone: "+123123123"
    projects:
      exampleProject: #TODO: gombio.batory.io/exampleProject - treat as slac channels
        healthChecks: [] #TODO:
        onCallSchedule: #TODO:
          - person: john@example.com
            start: "2017-10-10 13:00:00"
            end: "2017-10-10 14:00:00"
          - person: john@example.com
            start: "2017-10-10 16:00:00"
            end: "2017-10-10 20:00:00"
