// Copyright (c) 2013 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
btcd is a full-node bitcoin implementation written in Go.

The default options are sane for most users.  This means btcd will work 'out of
the box' for most users.  However, there are also a wide variety of flags that
can be used to control it.

The following section provides a usage overview which enumerates the flags.  An
interesting point to note is that the long form of all of these options
(except -C) can be specified in a configuration file that is automatically
parsed when btcd starts up.  By default, the configuration file is located at
~/.btcd/btcd.conf on POSIX-style operating systems and %APPDATA%\btcd\btcd.conf
on Windows.  The -C (--configfile) flag, as shown below, can be used to override
this location.

Usage:
  btcd [OPTIONS]

The flags are:
  -h, --help           Show this help message
  -V, --version        Display version information and exit
  -C, --configfile=    Path to configuration file
  -b, --datadir=       Directory to store data
  -a, --addpeer=       Add a peer to connect with at startup
      --connect=       Connect only to the specified peers at startup
      --nolisten       Disable listening for incoming connections -- NOTE:
                       Listening is automatically disabled if the --connect
                       option is used or if the --proxy option is used without
                       the --tor option
  -p, --port=          Listen for connections on this port (default: 8333,
                       testnet: 18333)
      --maxpeers=      Max number of inbound and outbound peers
      --banduration=   How long to ban misbehaving peers.  Valid time units are
                       {s, m, h}.  Minimum 1 second
  -u, --rpcuser=       Username for RPC connections
  -P, --rpcpass=       Password for RPC connections
  -r, --rpcport=       Listen for JSON/RPC messages on this port
      --norpc          Disable built-in RPC server -- NOTE: The RPC server is
                       disabled by default if no rpcuser/rpcpass is specified
      --nodnsseed      Disable DNS seeding for peers
      --proxy=         Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)
      --proxyuser=     Username for proxy server
      --proxypass=     Password for proxy server
      --tor            Specifies the proxy server used is a Tor node
      --testnet        Use the test network
      --regtest        Use the regression test network
      --nocheckpoints  Disable built-in checkpoints.  Don't do this unless you
                       know what you're doing
      --dbtype=        Database backend to use for the Block Chain
      --profile=       Enable HTTP profiling on given port -- NOTE port must be
                       between 1024 and 65536
  -d, --debuglevel=    Logging level {trace, debug, info, warn, error,
                       critical}

*/
package main
