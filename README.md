# platform

<p align="center">
    <img src="platform-hack/asset/platform-icon-256.png" width="256" height="256" alt="platform" />
</p>

Платформа для разработки, отладки, тестирования и развёртывания микросервисов.

## Документация

Запуск контейнера с документацией для локального просмотра:

```shell
docker run \
    --pull always \
    --rm \
    -it \
    -d \
    -p 8000:8000 \
    -v ${PWD}/platform-doc:/docs \
    squidfunk/mkdocs-material
```

Для оформления документации используется
соглашение [Semantic Line Breaks](https://sembr.org/).
