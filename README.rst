The Katzenpost Mix Network Daemons
==================================

Prerequisites
-------------

Building Katzenpost has the following prerequisites:

 * Some familiarity with building Go binaries.
 * `Go <https://golang.org>`_ 1.9 or later.
 * A recent version of `dep <https://github.com/golang/dep>`_.

Building
--------

.. code:: bash

    #
    # Fetch the Katzenpost components.
    #
    # Notes:
    #  * As of right now, due to the pace of development, the Katzenpost
    #    components are deliberately not vendored.
    #
    #  * This step SHOULD be omitted when doing development on Katzenpost,
    #    in favor of checking out all of the components into your GOPATH.
    #
    #  * The `daemons/vendor` directory contains all external dependencies
    #    that are required to build any part of Katzenpost, and can be
    #    reused if building things as libraries (Eg: symlinking
    #    `minclient/vendor` to `daemons/vendor` ).
    #
    dep ensure

    #
    # Build the binaries.
    #
    (cd authority/nonvoting; go build)
    (cd server; go build)
    (cd mailproxy; go build)

License
-------

`AGPL <https://www.gnu.org/licenses/agpl-3.0.en.html>`_: see the LICENSE file
for details.  The external dependencies have their own licenses.
