import pytest

from e2b import Sandbox, ProcessExitException


def test_kill_process(sandbox: Sandbox):
    cmd = sandbox.commands.run("sleep 10", background=True)
    pid = cmd.pid

    sandbox.commands.kill(pid)

    with pytest.raises(ProcessExitException):
        sandbox.commands.run(f"kill -0 {pid}")


def test_kill_non_existing_process(sandbox):
    non_existing_pid = 999999

    assert sandbox.commands.kill(non_existing_pid) == False
