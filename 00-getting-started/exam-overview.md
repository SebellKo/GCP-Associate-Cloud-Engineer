# 시험 안내와 공식 범위 매핑

확인일: 2026-11-02. 대상은 최초 취득용 **standard exam**이며 renewal과 구분한다.

| 항목 | 공식 안내 기준 |
|---|---|
| 시간 | 2시간 |
| 형식 | 객관식·복수 선택 50~60문항 |
| 응시료 | US$125, 해당 시 세금 별도 |
| 시험 언어 | 영어·일본어·스페인어·포르투갈어 |
| 사전 필수 자격 | 없음 |
| 권장 경험 | GCP 실습 6개월 이상 |
| 유효기간 | 3년 |

## 기준 문서 선택

[공식 자격증 페이지](https://cloud.google.com/learn/certification/cloud-engineer)의 Standard Exam Guide 링크가 연결하는 [PDF](https://services.google.com/fh/files/misc/associate_cloud_engineer_exam_guide_english.pdf)를 기준으로 했다. 확인 당시 PDF는 아래 4영역이다. 별도 [HTML 가이드](https://cloud.google.com/learn/certification/guides/cloud-engineer)는 5영역과 다른 비중을 표시했다. 이 교재는 그 차이를 숨기지 않고 standard 페이지에서 직접 연결된 PDF를 우선한다. 시험 예약 시에는 현재 공식 페이지가 연결하는 가이드를 다시 확인한다. 단순 확인일만으로 향후 시행일을 추정하지 않는다.

| PDF 영역 | 비중 | 챕터 | 모의고사당 문항 |
|---|---:|---|---:|
| 1. 환경 설정 | 약 20% | 03~06, 25 | 12 |
| 2. 계획 및 구현 | 약 30% | 07~20 | 18 |
| 3. 성공적 운영 | 약 30% | 07~14, 17~19, 21~23, 26 | 18 |
| 4. 접근과 보안 | 약 20% | 03~04, 24~26 | 12 |

## 세부 범위 확인표

| 범위 | 읽고 수행할 내용 |
|---|---|
| 1.1 | 03 계층·Cloud Identity·standalone 조직, 04 CLI, 05 API·quota·위치·Asset Inventory·AI 분석, 22 관측 설정, 25 Workforce federation |
| 1.2 | 06 결제 연결·예산·BigQuery export |
| 2.1 | 07 VM·OS Login·VM Manager, 08 디스크, 09 MIG·Spot, 16~19 Run·functions·GKE, 23 가속기·Agent Runtime |
| 2.2 | 10 객체·파일, 11 SQL·Spanner·AlloyDB, 12 문서·wide-column·분석·메시지·파이프라인 |
| 2.3 | 13 VPC·NGFW, 14 NAT·DNS·LB·tiers, 15 Shared VPC·peering·VPN·Interconnect |
| 2.4 | 20 Terraform·Config Connector·Helm·Fabric FAST·AI 지원 도구 |
| 3.1 | 07~09 VM·복구, 16~19 registry·Run·GKE 운영, 23 GPU/TPU·notebook·개발 환경 |
| 3.2 | 10~12 데이터 보안·수명 주기·백업·query/job·Database Center, 26 CMEK |
| 3.3 | 13~15 주소 범위·고정 IP·route·NAT·DNS·firewall |
| 3.4 | 21~23 로그·metric·agent·Prometheus·진단·서비스 상태·AI 지원 운영 |
| 4.1 | 03 상속, 24 IAM·역할·정책 |
| 4.2 | 25 서비스 계정·단기 자격 증명·impersonation·GKE/외부 federation |

고비용 제품과 조직 관리자 전용 기능은 개념·선택 기준·설계로 시작한다. 이 교재가 모든 제품의 프로덕션 운영 매뉴얼을 대체하지는 않는다. 실제 작업에는 해당 제품 문서의 전제 조건과 제약을 함께 확인한다.
