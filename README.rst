The Katzenpost Mix Network Daemons
==================================


Notes
-----

* The `daemons/vendor` directory contains all external dependencies
  that are required to build any part of Katzenpost, and can be
  reused if building things as libraries (Eg: symlinking
  `minclient/vendor` to `daemons/vendor` ).

* To build a new release see our release checklist document in our `docs`
  repository for the full release instructions.

* For a development work flow the Gopkg.toml may be edited to use the master
  branch of each Katzenpost repository.

License
-------

`AGPL <https://www.gnu.org/licenses/agpl-3.0.en.html>`_: see the LICENSE file
for details.  The external dependencies have their own licenses.
