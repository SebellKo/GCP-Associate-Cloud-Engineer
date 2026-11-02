# 모의고사 1 — 60문항

제한 시간 120분. 모든 문제는 자체 제작 학습 예제이며 실제 시험 문제나 공식 샘플의 복제가 아니다. 단일 선택은 한 개, 복수 선택은 정확히 두 개를 고른다. 정답 집합이 모두 일치하면 문항당 1점이며 부분 점수는 없다. 51/60은 자체 목표일 뿐 공식 합격선이 아니다.

영역 1: 12문항 · 영역 2: 18문항 · 영역 3: 18문항 · 영역 4: 12문항. 두 회는 서로 다른 상황·표현으로 복습하며 공식 난이도나 합격 확률을 측정하는 표준화 검사는 아니다.

답을 별도로 기록하고 다 푼 뒤 [해설](mock-exam-01-answers.md)을 연다. [오답 노트](error-log.md) · [전체 목차](../README.md)

## Q01. 단일 선택

조직 공통 위치 제한을 유지하면서 부서별 자원을 분리한다. 적절한 구조는?

A. 상위 위치 정책과 부서별 프로젝트
B. 모든 자원을 공개한다
C. 모든 직원에게 조직 Owner를 준다
D. 부서 이름만 다른 VM을 한 프로젝트에 둔다

## Q02. 단일 선택

새 팀원에게 같은 역할을 반복 부여하는 수고를 줄이려 한다.

A. 계정 비밀번호를 공유한다
B. Owner 키를 배포한다
C. 그룹에 역할을 부여하고 구성원을 관리한다
D. VM을 한 대 더 만든다

## Q03. 단일 선택

프로젝트에서 Compute API가 비활성화되어 생성 요청이 실패했다.

A. 해당 API 활성화 권한과 상태를 확인한다
B. 버킷을 공개한다
C. Cloud DNS를 재생성한다
D. VPC를 전부 삭제한다

## Q04. 단일 선택

리전 CPU 사용량이 quota 한도에 도달했다. 요구사항상 같은 리전이 필수다.

A. 필요 quota를 산정해 증액을 요청하고 승인 여부를 확인한다
B. 로그를 삭제한다
C. IAM 역할 이름을 바꾼다
D. 존 이름만 바꾸면 리전 한도가 항상 사라진다

## Q05. 단일 선택

같은 제품을 서울에 배포하려 한다. 구현 전 먼저 확인할 것은?

A. 버킷 이름 길이만
B. 로컬 브라우저 테마
C. 개발자 모니터 해상도
D. 제품·기능의 해당 리전 지원과 quota

## Q06. 단일 선택

학습자가 Cloud Shell에서 운영 프로젝트를 실수로 선택했다. 변경 전 대응은?

A. 로그만 지운다
B. Owner 권한을 더 준다
C. 활성 계정·프로젝트를 확인하고 학습 프로젝트를 명시한다
D. 그대로 apply한다

## Q07. 단일 선택

여러 프로젝트의 자산을 서비스별로 검색해야 한다.

A. 로컬 휴지통을 본다
B. 프로젝트별 결제 통화를 바꾼다
C. Cloud Asset Inventory를 활용한다
D. Cloud DNS 레코드만 본다

## Q08. 단일 선택

신입에게 프로젝트 Editor를 주었지만 결제 계정 사용 권한은 없다. 결제 연결 실패의 점검은?

A. VM을 Spot으로 바꾼다
B. 모든 디스크를 확장한다
C. 프로젝트 측·결제 계정 측 필요한 권한을 각각 확인한다
D. 서비스 계정 키를 공개한다

## Q09. 단일 선택

서비스별 월 비용을 SQL 보고서로 만들려 한다.

A. 로그 샘플만 보고 총액을 단정한다
B. NAT를 만든다
C. Billing export를 BigQuery에 설정한다
D. VM 이미지를 공개한다

## Q10. 복수 선택: 두 개

alerts-only 예산과 비용 관리를 설명한 것 중 맞는 두 가지를 고른다.

A. 모든 비용이 정확히 임계값에서 중단된다
B. 일반 알림형 예산은 자동 사용 차단이 아니다
C. 비용 보고·알림에는 지연이 있을 수 있다
D. 예산 삭제가 모든 자원을 삭제한다
E. VM 정지로 모든 연결 저장 비용도 사라진다

## Q11. 단일 선택

외부 기업 IdP 계정으로 직원의 콘솔 SSO를 구성하려 한다.

A. MIG를 만든다
B. Workload용 앱 키를 모든 직원에게 공유한다
C. Cloud NAT를 인증 서버로 쓴다
D. Workforce Identity Federation을 검토한다

## Q12. 단일 선택

Gemini가 자산 분석을 바탕으로 인프라 변경을 제안했다. 적용 전 조치는?

A. 현재 자원·권한·정책과 변경 계획을 검증한다
B. 권한을 모두 Owner로 바꾼다
C. AI가 제안했으므로 검토 없이 적용한다
D. 감사 로그를 끈다

## Q13. 단일 선택

기존 소프트웨어가 특정 OS 설정과 커널 모듈을 요구한다.

A. Firestore로 프로세스를 실행한다
B. Compute Engine을 검토한다
C. Cloud DNS를 쓴다
D. Cloud Run이 임의 커널 모듈을 허용한다고 가정한다

## Q14. 단일 선택

상태 없는 HTTP API이며 Kubernetes API는 필요 없고 운영 부담을 줄여야 한다.

A. Cloud Run을 검토한다
B. 전용 Kubernetes cluster를 반드시 직접 운영한다
C. 물리 서버를 구매한다
D. DB만 생성한다

## Q15. 단일 선택

Kubernetes 리소스와 제어가 필요하지만 노드 관리 부담은 줄인다.

A. Pub/Sub만 만든다
B. NAT에 앱 코드를 넣는다
C. Cloud Run이 모든 Kubernetes API를 노출한다고 생각한다
D. GKE Autopilot을 검토한다

## Q16. 단일 선택

중단 후 재시작 가능한 배치 작업의 VM 비용을 절감한다.

A. 모든 백업을 지운다
B. Owner를 준다
C. 단일 Spot이 절대 중단되지 않는다고 가정한다
D. Spot VM과 체크포인트·재시도를 사용한다

## Q17. 단일 선택

VM마다 CPU는 충분하지만 메모리가 부족하다. 비율을 조정하고 싶다.

A. 지원되는 custom machine type을 검토한다
B. DNS TTL만 조정한다
C. 방화벽 port만 늘린다
D. Nearline으로 바꾼다

## Q18. 단일 선택

읽기가 드문 파일을 1년 이상 보관하며 조회 비용도 감수할 수 있다.

A. 메모리 캐시만 원본으로 쓴다
B. Local SSD만 보관한다
C. 매일 조회가 많은 임시 파일과 똑같이 처리한다
D. Archive를 비용 후보로 비교한다

## Q19. 단일 선택

동일 리전에서 존 장애에 견디는 VM 그룹을 구성한다.

A. 모든 VM 이름만 다르게 한다
B. snapshot 없이 디스크를 지운다
C. 한 존에 모든 VM을 만든다
D. regional MIG를 여러 존에 배치한다

## Q20. 단일 선택

앱 여러 대가 공유 NFS 경로를 요구한다.

A. BigQuery를 NFS로 mount한다
B. Cloud DNS로 파일을 공유한다
C. Filestore를 비교한다
D. Pub/Sub topic에 파일 경로만 적는다

## Q21. 단일 선택

기존 PostgreSQL 앱을 엔진 호환성을 유지하며 관리형으로 옮긴다.

A. Cloud SQL for PostgreSQL을 우선 비교한다
B. Cloud Storage가 PostgreSQL 서버라고 간주한다
C. 모든 쿼리를 DNS로 보낸다
D. Pub/Sub를 SQL 엔진으로 쓴다

## Q22. 단일 선택

전 세계 사용자를 위한 관계형 트랜잭션에서 수평 확장과 강한 일관성이 필수다.

A. DNS TXT를 DB로 쓴다
B. Memorystore만 영구 원장으로 쓴다
C. Spanner를 검토한다
D. 객체 이름만 테이블처럼 쓴다

## Q23. 단일 선택

대량 과거 로그에 SQL 분석을 실행한다.

A. 모든 로그를 한 VM 메모리에 적재한다
B. readiness probe를 DB로 쓴다
C. BigQuery를 검토한다
D. Cloud NAT를 쿼리 엔진으로 쓴다

## Q24. 단일 선택

밀리초급 키 기반 대규모 시계열 접근이 중요하다.

A. Cloud DNS로 모든 측정값을 저장한다
B. Cloud Build를 DB로 쓴다
C. NAT 상태에 값을 저장한다
D. Bigtable을 후보로 비교한다

## Q25. 단일 선택

이벤트를 전달하는 생산자와 소비자를 분리하고 버퍼링한다.

A. Pub/Sub를 사용한다
B. Cloud SQL HA만 켠다
C. VM 이미지를 공개한다
D. 디스크 snapshot만 만든다

## Q26. 단일 선택

기존 VPC끼리 사설 IP로 연결하되 네트워크는 독립 관리한다.

A. Storage class를 같게 한다
B. project 이름만 같게 한다
C. DNS zone만 공유하면 경로가 자동 생긴다
D. CIDR 제약을 확인하고 VPC Peering을 검토한다

## Q27. 단일 선택

TCP/UDP 기반 내부 서비스가 필요하며 HTTP 경로 분기는 필요 없다.

A. Lifecycle 규칙을 만든다
B. 요구에 맞는 internal Network Load Balancer를 검토한다
C. Cloud DNS만으로 연결을 종료한다
D. Cloud NAT만 inbound 엔드포인트로 쓴다

## Q28. 단일 선택

회사 데이터센터를 IPsec으로 빠르게 연결한다.

A. Cloud VPN을 검토한다
B. Monitoring 경보만 만든다
C. OS Login만 켠다
D. Cloud Storage만 만든다

## Q29. 복수 선택: 두 개

IaC 적용 전에 수행할 올바른 검토 두 가지를 고른다.

A. state를 공개 Git에 올린다
B. plan이 실제 자원을 항상 변경한다고 생각한다
C. AI 생성 코드는 검토를 생략한다
D. plan의 생성·변경·삭제 대상을 확인한다
E. state·plan의 민감 정보를 보호한다

## Q30. 단일 선택

새 GKE 앱 이미지 경로를 버전 변경으로부터 정확히 고정한다.

A. 로그를 지운다
B. 검증한 image digest를 사용한다
C. latest tag만 사용한다
D. namespace 이름만 바꾼다

## Q31. 단일 선택

VM은 RUNNING이나 웹 요청이 실패한다. 우선 점검은?

A. 조직을 삭제한다
B. 상태가 RUNNING이므로 정상 처리한다
C. 결제 계정을 바꾼다
D. 앱 프로세스·포트·로그·네트워크를 확인한다

## Q32. 단일 선택

OS 표준 구성을 신규 VM에 반복 적용하려 한다.

A. 매번 무작위 이미지를 선택한다
B. DNS TTL을 늘린다
C. 검증된 custom image로 생성한다
D. 버킷을 공개한다

## Q33. 단일 선택

밤마다 VM 디스크의 복구 지점을 남긴다.

A. snapshot schedule을 검토한다
B. MIG 이름만 바꾼다
C. Public DNS zone을 만든다
D. CPU quota만 올린다

## Q34. 단일 선택

Cloud Run 새 버전이 오류를 내며 이전 버전은 정상이다.

A. 서비스를 무조건 공개한다
B. 로그 수집을 끈다
C. 모든 데이터를 삭제한다
D. DB 호환성 등을 확인하고 이전 revision으로 트래픽을 되돌린다

## Q35. 단일 선택

GKE Pod가 Pending이며 requests를 충족하는 capacity가 부족하다.

A. 버킷 등급을 변경한다
B. 스케줄링 이벤트·requests와 노드 capacity를 조사한다
C. 앱 로그만 무한 재조회한다
D. Cloud SQL backup을 복원한다

## Q36. 단일 선택

GKE Pod CPU에 따라 replica 수를 자동 조절한다.

A. snapshot을 예약한다
B. VPA만으로 replica 수를 바꾼다
C. HPA를 사용한다
D. Cloud DNS TTL을 늘린다

## Q37. 단일 선택

Cloud Storage object를 90일 후 등급 전환한다.

A. 모든 객체를 공개한다
B. VM을 autoscale한다
C. Lifecycle 조건과 action을 설정한다
D. 버킷 표시 이름만 바꾼다

## Q38. 단일 선택

Cloud SQL에서 잘못된 DELETE 직전으로 복구한다.

A. 인증을 끈다
B. CPU만 올린다
C. 백업/PITR 설정과 복구 가능 시점을 확인한다
D. HA standby가 과거 모든 시점을 갖는다고 가정한다

## Q39. 단일 선택

BigQuery job이 실패했다. 조사할 정보는?

A. job ID·location·권한·오류 내용
B. 사용자의 화면 밝기
C. VM hostname만
D. DNS TTL만

## Q40. 단일 선택

서비스 내부 도메인은 특정 VPC에서만 해석해야 한다.

A. VM 이름만 숨긴다
B. allUsers를 준다
C. public zone에 모두 게시한다
D. 해당 VPC가 허용된 private DNS zone을 사용한다

## Q41. 단일 선택

private VM의 Google API 접근이 필요하다. 인터넷 전체 발신과 구분하려 한다.

A. 디스크를 확장한다
B. Private Google Access 등 적절한 API 접근 경로를 검토한다
C. 모든 ingress를 연다
D. DNS를 무조건 지운다

## Q42. 단일 선택

오류 로그만 최근 30분 범위로 찾는다.

A. 예산 금액만 바꾼다
B. Logging에서 severity와 timestamp·resource 필터를 사용한다
C. 프로젝트를 삭제한다
D. VM 이미지만 복사한다

## Q43. 단일 선택

sink에 새 로그가 전달되지 않는다.

A. 필터와 writer identity의 목적지 쓰기 권한을 확인한다
B. 예산 이메일만 삭제한다
C. 디스크를 포맷한다
D. 모든 자원을 재시작한다

## Q44. 복수 선택: 두 개

VM의 guest metric과 경보를 운영할 때 맞는 두 가지를 고른다.

A. label cardinality는 비용에 영향이 없다
B. 필요한 guest metric은 Ops Agent 수집을 구성한다
C. notification channel의 전달을 확인한다
D. 경보 정책은 항상 VM 수를 늘린다
E. 기본 CPU metric이 모든 앱 로그도 보장한다

## Q45. 단일 선택

경보 조건은 충족됐지만 휴대폰 알림을 못 받는다.

A. NAT를 모두 삭제한다
B. notification channel과 검증 상태를 확인한다
C. DB를 복원한다
D. metric 이름만 지운다

## Q46. 단일 선택

요청 하나가 여러 서비스 중 어디에서 느린지 분석한다.

A. 모든 객체를 Archive로 바꾼다
B. Cloud DNS만 변경한다
C. 분산 trace를 검토한다
D. CPU 평균 하나만 보고 원인을 단정한다

## Q47. 단일 선택

조직 전체의 공통 서비스 오류가 의심된다.

A. 감사 로그를 끈다
B. Service Health의 관련 이벤트와 앱 증거를 비교한다
C. secret을 공개한다
D. 모든 VM을 무조건 삭제한다

## Q48. 단일 선택

Notebook 세션을 마쳤다. 비용 점검은?

A. Owner 권한만 제거하면 모든 자원이 삭제된다고 생각한다
B. 브라우저를 닫으면 모두 무료라고 가정한다
C. 실행 인스턴스·가속기와 연결 스토리지를 확인한다
D. 커널 출력만 지운다

## Q49. 단일 선택

앱이 버킷 하나의 object 읽기만 필요하다.

A. 해당 bucket에 objectViewer를 부여한다
B. 모든 bucket을 공개한다
C. 프로젝트 Editor를 준다
D. 조직 Owner를 준다

## Q50. 단일 선택

조직에서 부여한 Viewer를 프로젝트에서 제거하려 한다.

A. VM을 재시작한다
B. DNS를 바꾼다
C. 상위 binding과 그룹·유효 권한을 검토한다
D. 하위에 같은 binding이 없으면 무조건 접근 불가라 단정한다

## Q51. 단일 선택

새 직원에게 VM에 runtime 계정을 연결할 actAs 권한이 필요하다.

A. 적절한 범위의 Service Account User를 검토한다
B. 예산 조회 역할만 준다
C. Storage Object Viewer만 준다
D. DNS 관리자만 준다

## Q52. 단일 선택

계정을 impersonate하여 access token을 생성한다.

A. DB 테이블 이름만 바꾼다
B. VM network tag만 추가한다
C. 해당 서비스 계정의 Token Creator 권한을 검토한다
D. 버킷 등급만 바꾼다

## Q53. 복수 선택: 두 개

외부 CI에 장기 키 없이 접근을 주려 한다. 적절한 두 가지를 고른다.

A. Owner JSON key를 public Git에 둔다
B. Workload Identity Federation을 구성한다
C. Storage 공개로 인증을 대체한다
D. 외부 공급자의 모든 identity를 무조건 신뢰한다
E. issuer·audience·repo 조건으로 신뢰를 제한한다

## Q54. 단일 선택

GKE 앱마다 다른 Google API 권한을 갖게 한다.

A. 이미지 태그만 다르게 한다
B. 같은 Owner 키를 Secret에 복사한다
C. GKE Workload Identity Federation으로 Pod identity를 분리한다
D. 모든 앱에 노드 관리자 키를 공유한다

## Q55. 단일 선택

외부 OIDC 신뢰를 구성할 때 특정 repo의 승인된 작업만 허용한다.

A. 모든 사람에게 Owner를 준다
B. issuer·audience·attribute condition을 제한한다
C. 같은 공급자의 모든 identity를 무조건 신뢰한다
D. NAT를 인증 대신 쓴다

## Q56. 단일 선택

비밀이 Docker 이미지 layer에 포함됐다.

A. 태그만 바꾸고 그대로 둔다
B. 비밀을 회전하고 안전한 런타임 전달로 변경한다
C. README에 비밀을 추가한다
D. 이미지를 더 공개한다

## Q57. 단일 선택

KMS key가 비활성화되어 CMEK 데이터 접근에 문제가 생겼다.

A. CPU만 늘린다
B. DNS TTL만 줄인다
C. 데이터 IAM만 공개한다
D. 키 상태·의존 서비스·복구 영향을 확인한다

## Q58. 단일 선택

커스텀 IAM role을 만들기 전 우선 확인할 것은?

A. DNS zone 이름만
B. 모든 사용자 비밀번호
C. 편집기 글꼴
D. 적절한 predefined role과 필요한 permission 조합

## Q59. 단일 선택

서비스 계정의 자원 접근과 그 계정을 impersonate할 권한이 혼동된다.

A. 둘이 항상 같다고 간주한다
B. 자원 정책과 서비스 계정 자체 정책을 구분한다
C. 클러스터를 공개한다
D. 모든 계정을 지운다

## Q60. 단일 선택

새 앱에 최소 권한을 부여했는데 403이 발생한다.

A. 바로 프로젝트 Owner로 영구 확대한다
B. 실제 principal·resource·permission과 정책 제약을 확인한다
C. 모든 감사를 끈다
D. public access를 강제한다
