
import datetime
import os
import time

from peewee import (
    PostgresqlDatabase, CharField, Model, ForeignKeyField, BooleanField,
    DateTimeField, OperationalError
)

db = PostgresqlDatabase(
    os.getenv('DB'),
    user=os.getenv('USER_DB'),
    password=os.getenv('PASSWORD_DB'),
    host=os.getenv('HOST_DB')
)


class BaseModel(Model):
    """A base model that will use our Postgresql database"""

    class Meta:
        database = db