# 2. Use Docker Desktop instead of desktop VMs

Date: 2021-11-05

## Status

Superceded by [6. Use minikube as a primary k8s local provider](0006-use-minikube-as-a-primary-k8s-local-provider.md)

## Context

В качестве виртуальной машины для запуска `minikube`
использовать `Docker Desktop`.

## Decision

`Docker Desktop` обладает следующими преимуществами:

* Видит ресурсы за `VPN`-соединениями, не требует проброса `NAT`, в отличие
  от `hyperkit`.
* Позволяет обращаться по `localhost` к `docker`. Это реализовано с помощью
  проброса сокета `/var/run/docker.sock`.
* Используется большинством разработчиков.
* Вероятно, в будуещм использование `Docker Desktop` позволит сделать систему
  кроссплатформенной, добавив поддержку хостов на `Windows`.

## Consequences

Установка `Docker Desktop` легко автоматизируется с помощью `Homebrew`. Тем не
менее, могут встретиться разработчики, предпочитающие просто скачивать
приложения, перетягивая их в `/Applications`.

Драйвер `hyperkit` позволяет не иметь `docker-engine` на хост-машине — он
изолируется внутри виртуальной машины с `minikube`.

Так же `hyperkit` позволяет использовать `nfs` для синхронизации домашнего
каталога. Это самый быстрый способ синхронизировать файлы с виртуальной машиной.

Механизмов, предоставленных `Docker Desktop` для синхронизации файловой системы
с виртуальной машиной достаточно для комфортной работы.

Нельзя не отметить иконку кита, висящую в таскбаре рядом с часами. Она там ни к
селу ни к городу, если не выражаться более открыто.
