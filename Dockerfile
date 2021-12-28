FROM centos:centos7.9.2009

COPY ginPlus /home/admin/ginPlus

RUN  mkdir -p /home/admin/logs/ && chmod 0777 /home/admin/logs/

CMD  /home/admin/ginPlus >> /home/admin/logs/out.log 2>&1