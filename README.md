# Promo Tables

## Quick start

```bash
docker-compose up --build
```

## Usage

```bash
❯ make help
build-docker         Build docker image
build-go             Build
clean                Clean all build resources
flush                Drop all tables
help                 What you see here..
init                 Init tables and insert mockdata
test-clean           Expires all test results
test                 For CI and Local test
```

## Tables

![Image](./docs/promotion.png)

| Name                    | Description                                                         | TODO                                             |
| ----------------------- | ------------------------------------------------------------------- | ------------------------------------------------ |
| product                 | Products for store service                                          |                                                  |
| promo_category          | All promotions created by operator, uses as additional category     |                                                  |
| promo_controller        | Promotion factory, creates one promotion in several forms           | Manage the 'subscribe condition' column in Redis |
| user_store_subscription | Records user-specific promotion controller subscription information |                                                  |
| user                    | all users                                                           |                                                  |
| promo_purchased_log     | Table created to record and search purchase history                 |                                                  |

[DrawIO File](./docs/promotion.drawio)
[Description](https://velog.io/@d3fau1t/프로모션-상품을-관리하기위한-테이블-설계)
