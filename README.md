# Google Cloud Associate Cloud Engineer 독학 교재

프론트엔드 개발 경험이 있고 서버·네트워크·클라우드를 처음 배우는 학습자를 위한 **12주 ACE 취득 과정**이다. Linux·네트워크 기초에서 시작해 배포, 운영, IAM, 통합 실습으로 이어진다.

기준일: **2026-11-02**. [공식 자격증 안내](https://cloud.google.com/learn/certification/cloud-engineer)에서 연결한 **Standard Exam Guide PDF**를 기준으로 구성했다. 공식 HTML 가이드와 PDF의 영역 구성이 서로 다를 수 있으므로 [최신성 확인 방법](00-getting-started/exam-overview.md)을 먼저 읽는다.

## 여기서 시작

1. [학습 방법과 12주 일정](00-getting-started/study-guide.md)
2. [시험 안내와 범위 매핑](00-getting-started/exam-overview.md)
3. [실습 준비·비용·정리 원칙](00-getting-started/cost-safety.md)
4. [Chapter 01](chapters/chapter-01.md)부터 순서대로 학습

각 장은 **개념 → 개념 확인 3문항 → 실습/설계 → 시험형 예제 3문항 → 정답·오답 해설 → 완료 체크**로 구성된다. 정답은 접힌 영역에 두었다. 28장에 확인·예제 168문항과 별도 60문항 모의고사 2회를 제공한다. 모든 문제는 학습용 자체 제작이며 실제 출제 문항을 재현한 자료가 아니다.

## 전체 목차

| 장 | 학습 내용 | 공식 PDF 범위 |
|---|---|---|
| 01 | [클라우드와 Linux의 첫걸음](chapters/chapter-01.md) | 기초 |
| 02 | [IP·DNS·HTTP와 네트워크 기초](chapters/chapter-02.md) | 기초 |
| 03 | [리소스 계층·프로젝트·정책 상속](chapters/chapter-03.md) | 1.1 / 4.1 |
| 04 | [Cloud Shell·gcloud·인증과 구성](chapters/chapter-04.md) | 1.1 / 4.2 |
| 05 | [API·할당량·위치와 자산 관리](chapters/chapter-05.md) | 1.1 |
| 06 | [결제·예산·비용 내보내기](chapters/chapter-06.md) | 1.2 |
| 07 | [Compute Engine VM 생성과 접속](chapters/chapter-07.md) | 2.1 / 3.1 |
| 08 | [VM 디스크·이미지·스냅샷·복구](chapters/chapter-08.md) | 2.1 / 3.1 |
| 09 | [인스턴스 템플릿·MIG·가용성·Spot](chapters/chapter-09.md) | 2.1 / 3.1 |
| 10 | [Cloud Storage·수명 주기·파일 스토리지](chapters/chapter-10.md) | 2.2 / 3.2 |
| 11 | [관계형 DB·Cloud SQL·Spanner·AlloyDB](chapters/chapter-11.md) | 2.2 / 3.2 |
| 12 | [Firestore·BigQuery·Bigtable·데이터 파이프라인](chapters/chapter-12.md) | 2.2 / 3.2 |
| 13 | [VPC·서브넷·라우팅·방화벽](chapters/chapter-13.md) | 2.3 / 3.3 |
| 14 | [Cloud NAT·DNS·고정 IP·로드밸런서](chapters/chapter-14.md) | 2.3 / 3.3 |
| 15 | [Shared VPC·Peering·VPN·Interconnect](chapters/chapter-15.md) | 2.3 |
| 16 | [컨테이너·Artifact Registry·빌드](chapters/chapter-16.md) | 2.1 / 3.1 |
| 17 | [Cloud Run 배포·인증·리비전](chapters/chapter-17.md) | 2.1 / 3.1 |
| 18 | [이벤트 기반 함수·Eventarc·비동기 처리](chapters/chapter-18.md) | 2.1 / 3.1 |
| 19 | [GKE·Kubernetes 배포와 운영](chapters/chapter-19.md) | 2.1 / 3.1 |
| 20 | [Terraform·Helm·AI 보조 인프라 도구](chapters/chapter-20.md) | 2.4 |
| 21 | [Cloud Logging·Audit Logs·로그 라우팅](chapters/chapter-21.md) | 3.4 |
| 22 | [Monitoring·경보·Ops Agent·Prometheus](chapters/chapter-22.md) | 3.4 |
| 23 | [장애 진단·복구·최적화·AI 운영](chapters/chapter-23.md) | 3.1 / 3.2 / 3.3 / 3.4 |
| 24 | [IAM 역할·정책·최소 권한](chapters/chapter-24.md) | 4.1 |
| 25 | [서비스 계정·Impersonation·Workload Federation](chapters/chapter-25.md) | 4.2 |
| 26 | [Secret Manager·CMEK·보안 운영](chapters/chapter-26.md) | 3.2 / 4.1 / 4.2 |
| 27 | [통합 프로젝트: 배포부터 운영·정리까지](chapters/chapter-27.md) | 1 / 2 / 3 / 4 |
| 28 | [시험 전략·영역별 복습·다음 단계](chapters/chapter-28.md) | 1 / 2 / 3 / 4 |

## 종합 평가와 실습

- [모의고사 1 — 60문항 / 120분](07-exam-preparation/mock-exam-01.md) · [해설](07-exam-preparation/mock-exam-01-answers.md)
- [모의고사 2 — 60문항 / 120분](07-exam-preparation/mock-exam-02.md) · [해설](07-exam-preparation/mock-exam-02-answers.md)
- [실습 목록](labs/README.md) · [용어집](glossary.md) · [공식 참고 자료](resources.md)
- [학습 진도 기록](progress.md) · [오답 노트](07-exam-preparation/error-log.md)

주당 8~12시간을 출발점으로 잡고 어렵다면 기간을 늘린다. 공식 권장 경험은 GCP 실습 6개월 이상이며 이 12주 계획이나 자체 점수는 합격을 보장하지 않는다. 실제 시험을 예약하기 전에 공식 범위·언어·응시료·운영 정책을 다시 확인한다.

실습은 본인의 학습 프로젝트에서 직접 실행하는 방식이다. 교재 작성 중 실제 유료 GCP 자원을 생성하거나 end-to-end 배포를 검증하지 않았다. 명령 실행 전 위치·권한·제품 가용성·가격을 확인한다.
