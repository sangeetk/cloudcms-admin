FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD cloudcms-admin cloudcms-admin
EXPOSE 8080
CMD ["/cloudcms-admin"]