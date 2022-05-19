from unittest import TestCase
from app import *

_HELLO_WORLD = 'Hello world!'

class TestHelloWorld(TestCase):
  def test_always_passes(self):
    res = hello_world()

    self.assertTrue(res == _HELLO_WORLD)
