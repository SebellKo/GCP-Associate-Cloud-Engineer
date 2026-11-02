# Chapter 12. Firestore·BigQuery·Bigtable·데이터 파이프라인

[← 이전 장](chapter-11.md) · [전체 목차](../README.md) · [다음 장 →](chapter-13.md)

공식 가이드 연결: **2.2 / 3.2** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 데이터 모델로 제품을 고른다
Firestore는 document/collection 모델, Bigtable은 대규모 wide-column·키 기반 저지연 접근, BigQuery는 분석 SQL에 적합하다. Memorystore는 인메모리 캐시 용도이며 원본 데이터의 내구성 설계를 대신하지 않는다. 각 제품의 트랜잭션·쿼리·인덱스 제약을 요구사항과 비교한다.

Pub/Sub는 메시지 전달, Dataflow는 Apache Beam 기반 배치·스트림 처리다. Managed Service for Apache Kafka는 Kafka 호환 요구를 검토할 때 비교한다. 메시징과 변환 실행 엔진을 혼동하지 않는다. Pub/Sub 전달은 중복을 고려해 소비자를 멱등적으로 설계한다.

BigQuery는 실행 전에 예상 처리량을 보고 필요한 컬럼만 선택하며 partition filter를 사용한다. LIMIT만으로 읽는 바이트가 항상 줄어들지는 않는다. 데이터 로드 위치와 데이터셋 위치 호환성을 확인한다. 작업 실패는 job ID·위치·권한·오류를 점검한다. Storage Transfer Service는 대량 객체 이동에 적합하다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. BigQuery와 Cloud SQL의 주된 차이는?

2. Pub/Sub와 Dataflow의 차이는?

3. LIMIT 10이 항상 스캔 비용을 10행분으로 만드는가?

## 3. 실습과 관찰

BigQuery Sandbox 또는 무료 한도·과금 설정을 확인하고 `SELECT 1 AS ok`를 실행한다. Jobs에서 job ID와 상태를 기록한다. 추가 데이터 스캔은 예상 처리량을 먼저 확인한다. Pub/Sub 실습을 한다면 임시 topic/subscription만 만들고 발행→수신→ack 후 둘 다 삭제한다. 유료 원천 데이터를 스캔하지 않는다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

매일 수 TB의 과거 판매 데이터를 SQL 집계한다.

A. BigQuery를 사용한다
B. Cloud DNS를 쿼리한다
C. Memorystore만 원장으로 쓴다
D. 모든 데이터를 브라우저 메모리에 올린다

### 문제 2 — 단일 선택

앱이 유연한 JSON 형태의 사용자 문서를 저장하고 조회한다.

A. 스냅샷만 만든다
B. Firestore를 검토한다
C. VM 이미지에 매번 사용자 정보를 굽는다
D. DNS TXT 레코드를 DB로 쓴다

### 문제 3 — 단일 선택

메시지에서 값을 변환하고 window별 집계를 해야 한다.

A. Pub/Sub 수신과 Dataflow 처리를 조합한다
B. IAM 역할 이름만 바꾼다
C. Pub/Sub 자체가 임의 Beam 코드를 실행한다고 가정한다
D. 버킷 등급만 바꾼다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 분석과 트랜잭션 운영 중심이다.

2. 메시지 전달과 처리 파이프라인이다.

3. 아니다.

### 시험형 예제

**1. A** — 대규모 분석 SQL은 BigQuery의 용도다. 캐시·DNS·브라우저는 요구에 맞지 않는다.

**2. B** — document 모델이 맞다. DNS·스냅샷·이미지는 앱 문서 DB 역할이 아니다.

**3. A** — 전달과 처리 엔진을 분리한다. Pub/Sub는 일반 코드 처리 엔진이 아니며 등급·역할 이름도 변환을 실행하지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/products/databases)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-11.md) · [다음 장 →](chapter-13.md)
