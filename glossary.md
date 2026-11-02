# 용어집

| 용어 | 뜻·구분 |
|---|---|
| Resource hierarchy | 조직→폴더→프로젝트→자원 관리 구조 |
| Principal | 사용자·그룹·서비스 계정 등 접근 주체 |
| Permission / Role / Binding | 개별 작업 권한 / 권한 묶음 / 주체·역할을 자원에 연결 |
| Authentication / Authorization | 주체 확인 / 작업 허용 판단 |
| Region / Zone | 지리적 리전 / 리전 내 장애 격리 단위 |
| Quota / Capacity | 설정된 사용 한도 / 실제 공급 가능한 자원 |
| Ingress / Egress | 자원으로 들어오는 / 자원에서 나가는 통신 |
| CIDR | IP 주소 범위의 prefix 표기 |
| DNS / NAT | 이름 해석 / 주소 변환 |
| IAP | identity 기반 접근 프록시, TCP forwarding 등 |
| OS Login | IAM을 이용한 Linux 로그인 관리 |
| MIG | 템플릿 기반 managed VM 그룹 |
| Autoscaling / Autohealing | 규모 조정 / 비정상 자원 교체 |
| Snapshot / Image | 디스크 복구 지점 / VM 생성용 OS·소프트웨어 출발점 |
| RPO / RTO | 허용 데이터 손실 시간 / 복구 목표 시간 |
| HA / Backup / Replica | 가용성 구성 / 복구 사본 / 데이터 복제본 |
| Object / Block / File | 객체 API / 디스크 블록 / 파일 시스템 접근 모델 |
| OLTP / OLAP | 트랜잭션 처리 / 분석 처리 |
| Pod / Deployment / Service | 실행 단위 / 원하는 replica·rollout / 안정된 접근점 |
| HPA / VPA | Pod 수 확장 / Pod 자원 requests 조정 |
| Revision | 배포 설정·이미지의 특정 버전 |
| Concurrency | 한 실행 인스턴스에서 동시에 처리하는 요청 수 |
| Digest / Tag | 이미지 내용 식별자 / 재지정 가능한 이름 |
| ADC | 앱 라이브러리의 기본 자격 증명 탐색 방식 |
| Impersonation | 허용된 주체가 서비스 계정의 단기 자격 증명으로 작업 |
| Workload / Workforce Federation | 앱·자동화 identity / 사람 identity의 외부 신뢰 연결 |
| Secret / CMEK | 비밀 값 / 사용자가 관리하는 암호화 키 사용 방식 |
| Sink | 로그를 필터와 목적지로 라우팅하는 설정 |
| Metric / Log / Trace | 시계열 값 / 사건 기록 / 요청 경로·구간 |
| Cardinality | metric label 조합 등으로 생기는 시계열 수 |
| SLI / SLO | 서비스 지표 / 서비스 목표 |
| Least privilege | 필요한 작업과 자원 범위만 허용하는 최소 권한 |
| Operational overhead | 사람이 유지·운영해야 하는 부담 |
| Idempotency | 같은 요청·이벤트를 반복해도 의도한 업무 결과가 중복되지 않는 특성 |
| State / Plan | IaC가 관리하는 자원 연결 정보 / 적용 전 변경 계획 |
