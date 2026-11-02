# 모의고사 2 — 60문항

제한 시간 120분. 모든 문제는 자체 제작 학습 예제이며 실제 시험 문제나 공식 샘플의 복제가 아니다. 단일 선택은 한 개, 복수 선택은 정확히 두 개를 고른다. 정답 집합이 모두 일치하면 문항당 1점이며 부분 점수는 없다. 51/60은 자체 목표일 뿐 공식 합격선이 아니다.

영역 1: 12문항 · 영역 2: 18문항 · 영역 3: 18문항 · 영역 4: 12문항. 두 회는 서로 다른 상황·표현으로 복습하며 공식 난이도나 합격 확률을 측정하는 표준화 검사는 아니다.

답을 별도로 기록하고 다 푼 뒤 [해설](mock-exam-02-answers.md)을 연다. [오답 노트](error-log.md) · [전체 목차](../README.md)

## Q01. 단일 선택

회사의 dev/prod 비용·권한을 독립 관리하며 공통 정책을 유지한다.

A. 환경별 프로젝트와 상위 공통 정책을 사용한다
B. 모든 자원을 같은 공개 bucket에 둔다
C. 비밀번호를 공유한다
D. VM 이름 접두사만 다르게 한다

## Q02. 단일 선택

프로젝트 표시 이름을 바꿨는데 CLI가 이전 프로젝트 ID를 쓴다.

A. 방화벽을 연다
B. 표시 이름과 ID가 다른 식별자임을 이해하고 ID를 확인한다
C. DB를 삭제한다
D. 표시 이름 변경이 항상 ID를 바꾼다고 가정한다

## Q03. 단일 선택

새 입사자 20명에게 동일 업무 권한을 일괄 관리한다.

A. Cloud Identity 그룹과 그룹 binding을 사용한다
B. 각 VM에 같은 비밀번호를 넣는다
C. Owner JSON key를 공유한다
D. 모든 서비스를 공개한다

## Q04. 단일 선택

기본 gcloud configuration이 잘못되어 다른 프로젝트를 조회했다.

A. CLI 출력을 지운다
B. 예산을 없앤다
C. VM을 더 만든다
D. configuration을 분리하고 --project를 명시한다

## Q05. 복수 선택: 두 개

제품 배포 전에 확인할 전제로 맞는 두 가지를 고른다.

A. 위치 지원·quota·결제 상태를 확인한다
B. Owner만 있으면 모든 제품이 모든 존에 제공된다
C. API 활성화와 필요한 IAM 권한을 확인한다
D. 프로젝트 표시 이름만 바꾸면 API가 켜진다
E. quota가 충분하면 capacity도 항상 보장된다

## Q06. 단일 선택

특정 존에서 capacity 부족이지만 quota는 남아 있다.

A. 무조건 quota 오류라고 단정한다
B. Owner를 반복 부여한다
C. bucket 이름을 바꾼다
D. 허용되는 다른 존·머신 유형 또는 대기 전략을 검토한다

## Q07. 단일 선택

새 팀이 결제 계정의 비용을 모니터링해야 한다.

A. 프로젝트 VM admin이면 모든 결제를 본다고 가정한다
B. 필요한 결제 조회 권한을 별도 범위에서 검토한다
C. 사설 IP를 예약한다
D. image를 복사한다

## Q08. 단일 선택

alerts-only 예산을 설정한 뒤 실습 VM을 일주일 켰다.

A. 알림은 중단 장치가 아니므로 자원을 정리하고 비용을 확인한다
B. 이메일만 삭제한다
C. 디스크를 더 만든다
D. 예산보다 높게 과금될 수 없다고 단정한다

## Q09. 단일 선택

지난달부터 Billing export를 시작했다. 모든 이전 비용이 자동 소급될 거라 생각한다.

A. 모든 과거 데이터가 항상 복원된다고 보장한다
B. 지원되는 소급 범위를 확인하고 조기 export를 권장한다
C. DNS를 갱신한다
D. VM을 재부팅한다

## Q10. 단일 선택

조직에 있는 오래된 자원과 연결 정책을 검색한다.

A. 앱 memory를 늘린다
B. Asset Inventory에서 범위·필터·권한을 확인한다
C. DNS TTL만 조정한다
D. 콘솔 방문 기록만 본다

## Q11. 단일 선택

외부 인사 IdP를 사용하는 직원의 Cloud 접근을 연동한다.

A. bucket을 Archive로 바꾼다
B. Workforce Federation과 그룹·역할 매핑을 설계한다
C. VM service account 키를 모두 공유한다
D. Cloud NAT를 만든다

## Q12. 단일 선택

AI 도구가 리소스 목록을 요약했다. 운영 변경의 근거로 삼기 전 행동은?

A. 원본 인벤토리와 현재 정책·계획을 대조한다
B. 요약은 항상 최신·정확하다고 가정한다
C. 감사 기록을 없앤다
D. quota를 무조건 늘린다

## Q13. 단일 선택

VM의 무상태 웹 서버를 부하에 따라 늘리고 줄인다.

A. 스냅샷 이름만 바꾼다
B. private DNS zone만 만든다
C. 템플릿 기반 MIG autoscaling을 사용한다
D. bucket retention만 설정한다

## Q14. 단일 선택

새 VM 사양을 MIG에 rolling update로 배포한다.

A. 기존 템플릿을 임의로 덮어쓴다고 가정한다
B. 모든 이미지를 공개한다
C. 새 instance template을 만들고 그룹 버전을 갱신한다
D. DNS TTL만 늘린다

## Q15. 단일 선택

VM이 재생성돼도 남아야 하는 중요한 데이터를 저장한다.

A. 지원되는 영속 스토리지와 백업을 설계한다
B. 모든 로그만 남긴다
C. 컨테이너 임시 파일만 쓴다
D. Local SSD만 유일한 사본으로 쓴다

## Q16. 단일 선택

객체를 매일 읽는다. 비빈번 접근 등급의 조회 비용을 피하려 한다.

A. Coldline만 허용한다
B. Archive를 무조건 고른다
C. Standard를 검토한다
D. Nearline이 모든 읽기에 무료라 가정한다

## Q17. 단일 선택

문서 중심 모바일 앱의 데이터를 관리형으로 저장한다.

A. Firestore를 검토한다
B. NAT를 key-value DB로 쓴다
C. Cloud DNS TXT를 원장으로 쓴다
D. MIG를 문서 DB로 쓴다

## Q18. 단일 선택

PostgreSQL 호환 고성능 관리형 DB 제품을 비교한다.

A. Memorystore가 PostgreSQL 엔진이라고 간주한다
B. AlloyDB와 Cloud SQL의 적합성을 비교한다
C. Pub/Sub를 SQL 엔진으로 쓴다
D. Cloud Storage를 PostgreSQL로 mount한다

## Q19. 단일 선택

영구 원장이 별도로 있고 반복 읽기의 지연을 줄인다.

A. 캐시를 유일한 원장으로 바꾼다
B. NAT timeout만 낮춘다
C. Memorystore 같은 캐시를 검토한다
D. Cloud DNS에 모든 주문을 저장한다

## Q20. 단일 선택

Kafka 프로토콜 기반 기존 소비자를 관리형으로 연결한다.

A. Managed Service for Apache Kafka를 비교한다
B. VM snapshot만 만든다
C. DNS만 만든다
D. 모든 Kafka client가 수정 없이 Pub/Sub API를 쓴다고 가정한다

## Q21. 단일 선택

스트림 메시지에 window 집계와 변환을 적용한다.

A. bucket 등급 변경으로 계산한다
B. IAM role 이름으로 window를 만든다
C. Cloud DNS를 처리 엔진으로 쓴다
D. Dataflow 처리 파이프라인을 검토한다

## Q22. 단일 선택

부서별 프로젝트를 유지하며 네트워크팀만 VPC를 관리한다.

A. 프로젝트 이름만 통일한다
B. 모든 앱에 조직 관리자 권한을 준다
C. Shared VPC를 사용한다
D. Storage를 공개한다

## Q23. 복수 선택: 두 개

Interconnect와 네트워크 연결을 설계할 때 맞는 두 가지를 고른다.

A. DNS zone이 물리 회선을 대신한다
B. 암호화 요구를 별도로 확인한다
C. Interconnect는 모든 경우 자동 IPsec을 보장한다
D. 사설 경로만 있으면 모든 API IAM이 허용된다
E. 용량·위치·가용성 조건을 검토한다

## Q24. 단일 선택

두 VPC에 같은 CIDR이 있다. 바로 peering하면 충돌이 우려된다.

A. 겹침이 항상 무시된다고 가정한다
B. 주소 범위와 연결 제약을 재설계한다
C. MIG만 늘린다
D. image tag만 바꾼다

## Q25. 단일 선택

HTTP hostname과 path별로 요청을 라우팅한다.

A. Application Load Balancer를 검토한다
B. snapshot으로 분기한다
C. billing export로 트래픽을 보낸다
D. Cloud NAT로 L7 라우팅한다

## Q26. 단일 선택

내부 시스템이 같은 리전 여러 존에서 같은 subnet을 사용한다.

A. DNS가 IP 할당을 대체한다고 생각한다
B. regional subnet을 사용한다
C. subnet은 global이라고 가정한다
D. 각 존에 서로 겹치는 subnet을 만든다

## Q27. 단일 선택

앱 VM만 DB port에 접근하도록 한다.

A. source·target·port를 좁힌 ingress 규칙을 적용한다
B. DB를 공개 bucket으로 바꾼다
C. 모든 firewall을 지운다
D. 0.0.0.0/0 모든 port를 연다

## Q28. 단일 선택

Kubernetes 구성 패키지를 환경별 값으로 설치한다.

A. Billing export로 설치한다
B. Snapshot policy만 만든다
C. Helm을 검토한다
D. Cloud NAT를 package manager로 쓴다

## Q29. 단일 선택

팀에서 Terraform을 동시에 실행해 state 충돌이 우려된다.

A. 동시에 auto-approve한다
B. state를 매번 삭제한다
C. 잠금·접근 제어를 갖춘 remote backend를 설계한다
D. state를 public repo에 공유한다

## Q30. 단일 선택

Cloud Storage 객체 생성 이벤트로 함수를 실행한다.

A. 지원되는 Eventarc trigger와 대상·권한을 구성한다
B. VM hostname만 변경한다
C. MIG template만 복사한다
D. DNS zone만 추가한다

## Q31. 단일 선택

VM health check는 실패하지만 CPU 사용률은 낮다.

A. 권한을 모두 Owner로 바꾼다
B. 모든 데이터를 삭제한다
C. 앱 포트·응답 경로·방화벽·로그를 확인한다
D. CPU가 낮으므로 정상이라 단정한다

## Q32. 단일 선택

MIG VM에서 앱이 고장나면 자동으로 교체하려 한다.

A. snapshot 이름만 바꾼다
B. DNS zone만 만든다
C. autohealing과 적절한 health check를 구성한다
D. autoscaling만으로 모든 고장 교체가 보장된다고 본다

## Q33. 단일 선택

DB snapshot 복원 후 앱 데이터가 일관되지 않다.

A. 앱 일관성 백업 절차와 복원 검증을 설계한다
B. 로그를 지운다
C. snapshot은 언제나 앱 일관적이라 단정한다
D. CPU만 늘린다

## Q34. 단일 선택

Cloud Run 새 revision을 배포했지만 실제 사용자 트래픽은 보내지 않으려 한다.

A. DB부터 삭제한다
B. no-traffic 배포 후 테스트·트래픽 전환을 수행한다
C. 모든 버전을 즉시 100%로 보낸다
D. 로그를 끈다

## Q35. 단일 선택

Cloud Run DB 연결 수가 급증한다.

A. bucket class만 바꾼다
B. max instances가 모든 DB 연결을 정확히 보장한다고 단정한다
C. concurrency·인스턴스 수·connection pooling을 함께 검토한다
D. DNS TTL만 올린다

## Q36. 단일 선택

ImagePullBackOff인 Pod를 조사한다.

A. SQL backup을 복원한다
B. 로그 sink를 삭제한다
C. 이미지 경로·digest·pull identity 권한·네트워크를 확인한다
D. Pod memory만 올린다

## Q37. 단일 선택

Pod 한 개의 CPU·memory requests를 사용량에 맞춰 조정한다.

A. snapshot schedule을 만든다
B. VPA를 검토한다
C. DNS A record를 만든다
D. HPA가 단일 Pod requests만 바꾼다고 생각한다

## Q38. 단일 선택

Cloud SQL read replica를 만들었다. 자동 HA와 같은지 묻는다.

A. 모든 backup을 끈다
B. 읽기 복제와 HA standby 목적·failover를 구분한다
C. DB를 공개한다
D. 항상 완전히 동일하다고 설명한다

## Q39. 복수 선택: 두 개

BigQuery 비용 최적화에서 적절한 두 가지를 고른다.

A. 필요한 컬럼만 선택한다
B. DNS TTL이 쿼리 처리 바이트를 결정한다
C. LIMIT만으로 항상 반환 행수만 스캔한다
D. 권한을 Owner로 바꾸면 스캔 비용이 0이다
E. partition filter와 예상 처리량을 확인한다

## Q40. 단일 선택

Storage Transfer job이 실패했다.

A. CPU quota만 올린다
B. VM name만 바꾼다
C. source·destination 권한·위치·job 오류를 조사한다
D. 무조건 모든 객체를 공개한다

## Q41. 단일 선택

기존 subnet의 primary IPv4 범위를 늘려야 한다.

A. image tag를 바꾼다
B. 겹침·제약·기존 주소를 확인해 지원되는 확장을 수행한다
C. 같은 방식으로 언제나 축소도 가능하다고 가정한다
D. DNS name을 줄인다

## Q42. 단일 선택

NAT가 있는데 외부에서 private VM에 새 연결이 되지 않는다.

A. Public NAT는 발신용임을 확인하고 inbound 요구에 맞는 경로를 설계한다
B. SQL table 이름만 바꾼다
C. 객체 class만 바꾼다
D. NAT가 모든 inbound를 허용한다고 생각한다

## Q43. 단일 선택

방화벽 규칙 변경자를 알아야 한다.

A. 디스크 종류만 바꾼다
B. 관련 감사 로그에서 principal·method·timestamp를 조회한다
C. 이미지 크기만 비교한다
D. CPU metric만 본다

## Q44. 단일 선택

앱 사용자 ID를 metric label로 모두 넣어 시계열 수가 폭증했다.

A. 사용자마다 더 많은 label을 만든다
B. 비용은 영원히 동일하다고 가정한다
C. 모든 경보를 끈다
D. label cardinality를 줄이고 로그·metric 목적을 재설계한다

## Q45. 단일 선택

앱 CPU hotspot을 코드 함수 수준에서 찾는다.

A. DNS zone만 본다
B. billing account 이름만 바꾼다
C. Cloud Profiler 같은 profiling 도구를 검토한다
D. bucket retention을 늘린다

## Q46. 단일 선택

새 로그 sink를 만들고 3개월 전 로그가 자동 전송되길 기대한다.

A. sink가 모든 과거 로그를 무조건 소급한다고 생각한다
B. 신규 라우팅과 과거 데이터 복사 요구를 구분한다
C. quota만 늘린다
D. VM을 재부팅한다

## Q47. 단일 선택

가속기 notebook 사용이 끝났지만 세션은 계속 실행 중이다.

A. 비밀번호만 바꾼다
B. 셀 출력을 지운다
C. 인스턴스·GPU/TPU·저장소 수명을 확인하고 정리한다
D. 브라우저 탭만 닫는다

## Q48. 단일 선택

Active Assist가 운영 VM 축소를 권고했다.

A. 감사 로그를 삭제한다
B. DB 백업도 삭제한다
C. 피크 부하·주기 업무·SLO를 확인해 결정한다
D. 권고는 어떤 상황에서도 바로 적용한다

## Q49. 단일 선택

프로젝트의 모든 bucket 관리 권한이 있는 앱을 한 bucket 읽기로 축소한다.

A. bucket을 공개한다
B. 기존 광범위 binding을 검토·제거하고 필요한 bucket 읽기만 부여한다
C. 좁은 역할을 추가하면 상위 광범위 권한이 자동 취소된다고 생각한다
D. 계정 이름만 바꾼다

## Q50. 단일 선택

필요 작업을 predefined role이 정확히 담지 못한다.

A. 무조건 Owner를 준다
B. 지원 permission과 유지보수 부담을 확인해 custom role을 검토한다
C. 모든 리소스를 공개한다
D. 비밀번호를 공유한다

## Q51. 단일 선택

직원이 Cloud Run에 특정 runtime 계정을 연결해 배포한다.

A. DNS viewer만 준다
B. 배포 권한과 해당 서비스 계정 actAs 권한을 확인한다
C. 버킷 objectViewer만 준다
D. 노트북 화면 잠금만 푼다

## Q52. 단일 선택

impersonation은 성공했지만 bucket 읽기가 403이다.

A. DNS를 지운다
B. 서비스 계정의 해당 bucket 데이터 권한을 확인한다
C. CPU만 늘린다
D. token 생성 권한이면 모든 자원 접근도 보장된다고 생각한다

## Q53. 단일 선택

외부 CI가 모든 GitHub repo에서 federation하도록 설정됐다. 특정 repo만 허용해야 한다.

A. provider attribute condition과 principal binding 범위를 좁힌다
B. 모든 repo를 계속 신뢰한다
C. NAT를 늘린다
D. Owner를 전 세계에 준다

## Q54. 단일 선택

GKE 앱 A/B가 서로 다른 bucket을 읽어야 한다.

A. 이미지 이름만 다르게 한다
B. workload별 identity와 bucket 범위 역할을 분리한다
C. namespace 이름만으로 IAM이 자동 분리된다고 생각한다
D. 동일 Owner key를 두 앱에 복사한다

## Q55. 단일 선택

관리형 서비스의 service agent 역할이 낯설어 삭제하려 한다.

A. 이름만 보고 무조건 삭제한다
B. 모든 service agent를 개인 계정으로 바꾼다
C. 비밀을 로그에 출력한다
D. 해당 제품의 service agent 역할·의존성을 확인한다

## Q56. 단일 선택

앱이 Secret Manager의 비밀 하나만 읽어야 한다.

A. secret을 공개 Git에 넣는다
B. 프로젝트 Owner를 준다
C. 그 secret 범위의 필요한 secret 접근 권한을 준다
D. KMS 관리 역할만으로 모든 secret을 읽는다고 생각한다

## Q57. 단일 선택

CMEK를 도입하니 IAM을 생략해도 되는지 묻는다.

A. 암호화면 모든 principal의 읽기가 안전하다고 생각한다
B. 암호화 키 통제와 데이터 접근 권한을 함께 설계한다
C. 모든 bucket을 공개한다
D. backup을 없앤다

## Q58. 복수 선택: 두 개

서비스 계정 private key가 유출됐다. 적절한 두 가지를 고른다.

A. 다른 repository에 키를 복사한다
B. 키를 폐기·회전한다
C. 사용 흔적을 조사하고 키 없는 인증을 검토한다
D. Git 파일만 지우고 키를 계속 쓴다
E. 확장자를 바꾸면 키가 무효화된다

## Q59. 단일 선택

allow binding이 있지만 IAM condition의 시간 조건이 끝났다.

A. DNS만 바꾼다
B. 조건과 실제 요청 맥락을 검토한다
C. VM만 증설한다
D. allow이면 언제나 무조건 허용된다고 가정한다

## Q60. 단일 선택

같은 그룹에 속한 사용자의 역할을 제거했는데 접근이 남는다.

A. 브라우저를 닫으면 IAM이 제거된다고 생각한다
B. 디스크를 지운다
C. bucket 이름만 바꾼다
D. 직접·그룹·상속 등 남아 있는 부여 경로를 조사한다
