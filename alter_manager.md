```
global:
  resolve_timeout: 5m
  smtp_smarthost: 'smtp.qiye.aliyun.com:465'
  smtp_from: 'alertsyt@test.com'
  smtp_auth_username: 'alertsyt@test.com'
  smtp_auth_password: 'passwd'
  smtp_hello: 'alertsyt'
  smtp_require_tls: false
  
route:
  group_by: ['alertname']
  group_wait: 30s
  group_interval: 5m
  repeat_interval: 5m
  receiver: default

  routes:
  - receiver: email
    group_wait: 10s
receivers:
- name: 'default'
  email_configs:
  - to: 'email@test.com'
    send_resolved: true

- name: 'email'
  email_configs:
  - to: 'email@test.com'
    send_resolved: true
```
